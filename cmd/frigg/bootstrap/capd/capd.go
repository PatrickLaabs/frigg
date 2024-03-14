package capd

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/cluster"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/workloadcluster"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "capd",
		Short: "clusterapi provider docker",
		Long:  "Creates local Kubernetes clusters using clusterapi's provider capd (docker)",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("subcommand is required")
		},
	}
	c.AddCommand(cluster.NewCommand(logger, streams))
	c.AddCommand(workloadcluster.NewCommand())
	return c
}
