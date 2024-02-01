/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package export implements the `export` command
package export

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/export/kubeconfig"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/export/logs"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
)

// NewCommand returns a new cobra.Command for export
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "export",
		Short: "Exports one of [kubeconfig, logs]",
		Long:  "Exports one of [kubeconfig, logs]",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("subcommand is required")
		},
	}
	// add subcommands
	c.AddCommand(logs.NewCommand(logger, streams))
	c.AddCommand(kubeconfig.NewCommand(logger))
	return c
}
