package create_template

import (
	"github.com/PatrickLaabs/frigg/cmd"
	mgmt_cluster "github.com/PatrickLaabs/frigg/cmd/frigg/gitops/create-template/mgmt-cluster"
	workload_cluster "github.com/PatrickLaabs/frigg/cmd/frigg/gitops/create-template/workload-cluster"
	"github.com/PatrickLaabs/frigg/pkg/errors"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for get
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "create-template",
		Short: "gitops template generation",
		Long:  "gitops template generation",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("Subcommand is required")
		},
	}
	c.AddCommand(mgmt_cluster.NewCommand(logger, streams))
	c.AddCommand(workload_cluster.NewCommand())
	return c
}
