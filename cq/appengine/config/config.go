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

// package config implements validation and common manipulation of CQ config
// files.
package config

import (
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"go.chromium.org/luci/config/validation"

	v1 "go.chromium.org/luci/cq/api/config/v1"
	v2 "go.chromium.org/luci/cq/api/config/v2"
)

// Config validation rules go here.

func init() {
	validation.Rules.Add("regex:projects/[^/]+", "cq.cfg", validateProject)
	validation.Rules.Add("regex:projects/[^/]+/refs/.+", "cq.cfg", validateRef)
}

// validateRefCfg validates legacy ref-specific cq.cfg.
// Validation result is returned via validation ctx,
// while error returned directly implies only a bug in this code.
func validateRef(ctx *validation.Context, configSet, path string, content []byte) error {
	ctx.SetFile(path)
	cfg := v1.Config{}
	if err := proto.UnmarshalText(string(content), &cfg); err != nil {
		ctx.Error(err)
		return nil
	}
	validateV1(ctx, &cfg)
	return nil
}

// validateProjectCfg validates project-level cq.cfg.
// Validation result is returned via validation ctx,
// while error returned directly implies only a bug in this code.
func validateProject(ctx *validation.Context, configSet, path string, content []byte) error {
	ctx.SetFile(path)
	cfg := v2.Config{}
	if err := proto.UnmarshalText(string(content), &cfg); err != nil {
		ctx.Error(err)
	} else {
		validateProjectConfig(ctx, &cfg)
	}
	return nil
}

func validateProjectConfig(ctx *validation.Context, cfg *v2.Config) {
	if cfg.DrainingStartTime != "" {
		if _, err := time.Parse(time.RFC3339, cfg.DrainingStartTime); err != nil {
			ctx.Errorf("failed to parse draining_start_time %q as RFC3339 format: %s", cfg.DrainingStartTime, err)
		}
	}
	if cfg.CqStatusHost != "" {
		switch u, err := url.Parse("https://" + cfg.CqStatusHost); {
		case err != nil:
			ctx.Errorf("failed to parse cq_status_host %q: %s", cfg.CqStatusHost, err)
		case u.Host != cfg.CqStatusHost:
			ctx.Errorf("cq_status_host %q should be just a host %q", cfg.CqStatusHost, u.Host)
		}
	}
	if cfg.SubmitOptions != nil {
		ctx.Enter("submit_options")
		if cfg.SubmitOptions.MaxBurst < 0 {
			ctx.Errorf("max_burst must be >= 0")
		}
		if cfg.SubmitOptions.BurstDelay != nil {
			switch d, err := ptypes.Duration(cfg.SubmitOptions.BurstDelay); {
			case err != nil:
				ctx.Errorf("invalid burst_delay: %s", err)
			case d.Seconds() < 0.0:
				ctx.Errorf("burst_delay must be positive or 0")
			}
		}
		ctx.Exit()
	}
	if len(cfg.ConfigGroups) == 0 {
		ctx.Errorf("at least 1 config_group is required")
		return
	}
	for i, g := range cfg.ConfigGroups {
		ctx.Enter("config_group #%d", i+1)
		validateConfigGroup(ctx, g)
		ctx.Exit()
	}
}

func validateConfigGroup(ctx *validation.Context, g *v2.ConfigGroup) {
	// TODO(tandrii): implement.
}
