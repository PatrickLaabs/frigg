/*
Copyright © 2024 Patrick Laabs <patrick.laabs@me.com>
*/
package kind

import (
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd"
	createcluster "github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/create/cluster"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/errors"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
	"github.com/spf13/cobra"
)

// KindApiUsageCmd represents the localBootstrap command
//var KindApiUsageCmd = &cobra.Command{
//	Use:   "kindApiUsageCmd",
//	Short: "Using Go Pkg of KinD",
//	Long: `A longer description that spans multiple lines and likely contains examples
//and usage of using your command. For example:
//
//Cobra is a CLI library for Go that empowers applications.
//This application is a tool to generate the needed files
//to quickly create a Cobra application.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("KindApiUsageCmd called")
//	},
//}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "create",
		Short: "Creates one of [cluster]",
		Long:  "Creates one of local Kubernetes cluster (cluster)",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("Subcommand is required")
		},
	}
	cmd.AddCommand(createcluster.NewCommand(logger, streams))
	return cmd
}
