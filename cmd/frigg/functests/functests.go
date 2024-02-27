package functests

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/cmd/frigg/functests/getkubeconfig"
	"github.com/PatrickLaabs/frigg/cmd/frigg/functests/helmchartproxy"
	"github.com/PatrickLaabs/frigg/cmd/frigg/functests/mgmtgen"
	"github.com/PatrickLaabs/frigg/cmd/frigg/functests/modifykubeconfig"
	"github.com/PatrickLaabs/frigg/cmd/frigg/functests/reporender"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "functests",
		Short: "functests",
		Long:  "functests",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("subcommand is required")
		},
	}
	c.AddCommand(getkubeconfig.NewCommand(logger, streams))
	c.AddCommand(modifykubeconfig.NewCommand())
	c.AddCommand(mgmtgen.NewCommand(logger, streams))
	c.AddCommand(helmchartproxy.NewCommand())
	c.AddCommand(reporender.NewCommand(logger, streams))
	return c
}
