package createbootstrap

import (
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd"
	argohubcluster "github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/createbootstrap/cluster"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/errors"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for get
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "createbootstrap",
		Short: "Creates a clusters with gitops onboard",
		Long:  "Creates a local kind clusters, and deploys the argohub gitops platform on it",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("Subcommand is required")
		},
	}
	c.AddCommand(argohubcluster.NewCommand(logger, streams))
	return c
}
