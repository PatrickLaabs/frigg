package harvester

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/harvester/cluster"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// 01-Get Started
// The Cluster-API Provider for Harvester is in an early-early stage.

// 02-Upload an Image
// On your existing Harvester Instance, you should upload an Ubuntu Image.
// You can get from here: https://cloud-images.ubuntu.com/jammy/current/
// I choosed the `jammy-server-cloudimg-amd64.img` as it is recommended by the CAPHV Provider.

// 03-Generate a Public SSH Key
// this can be done inside your harvester dashboard at `Advanced > SSH Keys`
// This will later be used for your ClusterAPI CAPHV Manifest and Cluster creation.

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "harvester",
		Short: "harvester",
		Long:  "harvester",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("subcommand is required")
		},
	}
	c.AddCommand(cluster.NewCommand(logger, streams))
	return c
}
