// Copyright 2017 The LUCI Authors.
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

package cli

import (
	"github.com/maruel/subcommands"

	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/flag"

	"go.chromium.org/luci/machine-db/api/crimson/v1"
)

// printOSes prints operating system data to stdout in tab-separated columns.
func printOSes(showHeaders bool, oses ...*crimson.OS) {
	if len(oses) > 0 {
		p := newStdoutPrinter()
		defer p.Flush()
		if showHeaders {
			p.Row("Name", "Description")
		}
		for _, os := range oses {
			p.Row(os.Name, os.Description)
		}
	}
}

// GetOSesCmd is the command to get operating systems.
type GetOSesCmd struct {
	subcommands.CommandRunBase
	req crimson.ListOSesRequest
	f   FormattingFlags
}

// Run runs the command to get operating systems.
func (c *GetOSesCmd) Run(app subcommands.Application, args []string, env subcommands.Env) int {
	ctx := cli.GetContext(app, c, env)
	client := getClient(ctx)
	resp, err := client.ListOSes(ctx, &c.req)
	if err != nil {
		errors.Log(ctx, err)
		return 1
	}
	printOSes(c.f.showHeaders, resp.Oses...)
	return 0
}

// getOSesCmd returns a command to get operating systems.
func getOSesCmd() *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "get-oses [-name <name>]...",
		ShortDesc: "retrieves operating systems",
		LongDesc:  "Retrieves operating systems matching the given names, or all operating systems if names are omitted.",
		CommandRun: func() subcommands.CommandRun {
			cmd := &GetOSesCmd{}
			cmd.Flags.Var(flag.StringSlice(&cmd.req.Names), "name", "Name of an operating system to filter by. Can be specified multiple times.")
			cmd.f.Register(cmd)
			return cmd
		},
	}
}
