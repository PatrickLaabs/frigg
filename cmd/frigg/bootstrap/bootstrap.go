package bootstrap

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd"
	capd_controller "github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd-controller"
	"github.com/PatrickLaabs/frigg/pkg/errors"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for get
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "bootstrap",
		Short: "bootstrap various clusters on different providers",
		Long:  "bootstrap various clusters on different providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("Subcommand is required")
		},
	}
	c.AddCommand(capd.NewCommand(logger, streams))
	c.AddCommand(capd_controller.NewCommand(logger, streams))
	return c
}
