// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"strings"

	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"

	"go.chromium.org/luci/machine-db/api/crimson/v1"
	"go.chromium.org/luci/machine-db/appengine/database"
)

// CreateVM handles a request to create a new VM.
func (*Service) CreateVM(c context.Context, req *crimson.CreateVMRequest) (*crimson.VM, error) {
	if err := createVM(c, req.Vm); err != nil {
		return nil, err
	}
	return req.Vm, nil
}

// ListVMs handles a request to list VMs.
func (*Service) ListVMs(c context.Context, req *crimson.ListVMsRequest) (*crimson.ListVMsResponse, error) {
	vlans := make(map[int64]struct{}, len(req.Vlans))
	for _, vlan := range req.Vlans {
		vlans[vlan] = struct{}{}
	}
	vms, err := listVMs(c, stringset.NewFromSlice(req.Names...), vlans)
	if err != nil {
		return nil, internalError(c, err)
	}
	return &crimson.ListVMsResponse{
		Vms: vms,
	}, nil
}

// createVM creates a new VM in the database.
func createVM(c context.Context, v *crimson.VM) (err error) {
	if err := validateVMForCreation(v); err != nil {
		return err
	}
	tx, err := database.Begin(c)
	if err != nil {
		return internalError(c, errors.Annotate(err, "failed to begin transaction").Err())
	}
	defer func() {
		if err != nil {
			if e := tx.Rollback(); e != nil {
				errors.Log(c, errors.Annotate(e, "failed to roll back transaction").Err())
			}
		}
	}()

	// TODO(smut): Check that the provided IP address is unassigned.

	// By setting hostnames.vlan_id as both FOREIGN KEY and NOT NULL when setting up the database,
	// we can avoid checking if the given VLAN is valid. MySQL will ensure the given VLAN exists.
	res, err := tx.ExecContext(c, `
		INSERT INTO hostnames (name, vlan_id)
		VALUES (?, ?)
	`, v.Name, v.Vlan)
	if err != nil {
		switch e, ok := err.(*mysql.MySQLError); {
		case !ok:
			// Type assertion failed.
		case e.Number == mysqlerr.ER_DUP_ENTRY && strings.Contains(e.Message, "'name'"):
			// e.g. "Error 1062: Duplicate entry 'hostname-vlanId' for key 'name'".
			return status.Errorf(codes.AlreadyExists, "duplicate hostname %q for VLAN %d", v.Name, v.Vlan)
		case e.Number == mysqlerr.ER_NO_REFERENCED_ROW_2 && strings.Contains(e.Message, "`vlan_id`"):
			// e.g. "Error 1452: Cannot add or update a child row: a foreign key constraint fails (FOREIGN KEY (`vlan_id`) REFERENCES `vlans` (`id`))".
			return status.Errorf(codes.NotFound, "unknown VLAN %d", v.Vlan)
		}
		return internalError(c, errors.Annotate(err, "failed to create hostname").Err())
	}
	hostnameId, err := res.LastInsertId()
	if err != nil {
		return internalError(c, errors.Annotate(err, "failed to fetch hostname").Err())
	}

	// vms.hostname_id, vms.physical_host_id, and vms.os_id are NOT NULL as above.
	_, err = tx.ExecContext(c, `
		INSERT INTO vms (hostname_id, physical_host_id, os_id, description, deployment_ticket)
		VALUES (
			?,
			(SELECT p.id FROM physical_hosts p, hostnames h WHERE p.hostname_id = h.id AND h.name = ? AND h.vlan_id = ?),
			(SELECT id FROM oses WHERE name = ?),
			?,
			?
		)
	`, hostnameId, v.Host, v.Vlan, v.Os, v.Description, v.DeploymentTicket)
	if err != nil {
		switch e, ok := err.(*mysql.MySQLError); {
		case !ok:
			// Type assertion failed.
		case e.Number == mysqlerr.ER_BAD_NULL_ERROR && strings.Contains(e.Message, "'physical_host_id'"):
			// e.g. "Error 1048: Column 'physical_host_id' cannot be null".
			return status.Errorf(codes.NotFound, "unknown physical host %q for VLAN %d", v.Host, v.Vlan)
		case e.Number == mysqlerr.ER_BAD_NULL_ERROR && strings.Contains(e.Message, "'os_id'"):
			// e.g. "Error 1048: Column 'os_id' cannot be null".
			return status.Errorf(codes.NotFound, "unknown operating system %q", v.Os)
		}
		return internalError(c, errors.Annotate(err, "failed to create VM").Err())
	}

	// TODO(smut): Assign the provided IP address.

	if err := tx.Commit(); err != nil {
		return internalError(c, errors.Annotate(err, "failed to commit transaction").Err())
	}
	return nil
}

// listVMs returns a slice of VMs in the database.
func listVMs(c context.Context, names stringset.Set, vlans map[int64]struct{}) ([]*crimson.VM, error) {
	db := database.Get(c)
	rows, err := db.QueryContext(c, `
		SELECT hv.name, hv.vlan_id, hp.name, o.name, vm.description, vm.deployment_ticket
		FROM vms vm, hostnames hv, physical_hosts p, hostnames hp, oses o
		WHERE vm.hostname_id = hv.id
			AND vm.physical_host_id = p.id
			AND p.hostname_id = hp.id
			AND vm.os_id = o.id
	`)
	// TODO(smut): Fetch the assigned IP address.
	if err != nil {
		return nil, errors.Annotate(err, "failed to fetch VMs").Err()
	}
	defer rows.Close()

	var vms []*crimson.VM
	for rows.Next() {
		v := &crimson.VM{}
		if err = rows.Scan(&v.Name, &v.Vlan, &v.Host, &v.Os, &v.Description, &v.DeploymentTicket); err != nil {
			return nil, errors.Annotate(err, "failed to fetch VM").Err()
		}
		// TODO(smut): use the database to filter rather than fetching all entries.
		if _, ok := vlans[v.Vlan]; matches(v.Name, names) && (ok || len(vlans) == 0) {
			vms = append(vms, v)
		}
	}
	return vms, nil
}

// validateVMForCreation validates a VM for creation.
func validateVMForCreation(v *crimson.VM) error {
	switch {
	case v == nil:
		return status.Error(codes.InvalidArgument, "VM specification is required")
	case v.Name == "":
		return status.Error(codes.InvalidArgument, "hostname is required and must be non-empty")
	case v.Vlan < 1:
		return status.Error(codes.InvalidArgument, "VLAN is required and must be positive")
	case v.Host == "":
		return status.Error(codes.InvalidArgument, "physical hostname is required and must be non-empty")
	case v.Os == "":
		return status.Error(codes.InvalidArgument, "operating system is required and must be non-empty")
	default:
		return nil
	}
}