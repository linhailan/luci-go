// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package adminsrv implements Admin API.
//
// Code defined here is either invoked by an administrator or by the service
// itself (via cron jobs or task queues).
package adminsrv

import (
	"github.com/luci/luci-go/appengine/gaeauth/server/gaesigner"

	"github.com/luci/luci-go/tokenserver/appengine/certconfig"
	"github.com/luci/luci-go/tokenserver/appengine/machinetoken"

	"github.com/luci/luci-go/tokenserver/api/admin/v1"
)

// serverImpl implements admin.AdminServer RPC interface.
type serverImpl struct {
	certconfig.ImportCAConfigsRPC
	machinetoken.InspectMachineTokenRPC
}

// NewServer returns prod AdminServer implementation.
//
// It assumes authorization has happened already.
func NewServer() admin.AdminServer {
	return &serverImpl{
		InspectMachineTokenRPC: machinetoken.InspectMachineTokenRPC{
			Signer: gaesigner.Signer{},
		},
	}
}
