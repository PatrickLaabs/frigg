/*
Copyright 2019 The Kubernetes Authors.

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

// Package load implements the `load` command
package load

import (
	"errors"

	"github.com/spf13/cobra"

	dockerimage "github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/load/docker-image"
	imagearchive "github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/load/image-archive"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
)

// NewCommand returns a new cobra.Command for get
func NewCommand(logger log.Logger) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "load",
		Short: "Loads images into nodes",
		Long:  "Loads images into node from an archive or image on host",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("subcommand is required")
		},
	}
	// add subcommands
	c.AddCommand(dockerimage.NewCommand(logger))
	c.AddCommand(imagearchive.NewCommand(logger))
	return c
}
