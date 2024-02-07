package bootstrap

import (
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capv"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capz"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/errors"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for get
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "bootstrap",
		Short: "bootstrap",
		Long:  "bootstrap",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("Subcommand is required")
		},
	}
	c.AddCommand(capd.NewCommand(logger, streams))
	c.AddCommand(capv.NewCommand(logger, streams))
	c.AddCommand(capz.NewCommand(logger, streams))
	return c
}
