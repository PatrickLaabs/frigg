package capd

import (
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd/cluster"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "capd",
		Short: "capd",
		Long:  "capd",
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
