package getkubeconfig

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/pkg/common/kubeconfig"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "get kubeconfig func",
		Short: "get kubeconfig func",
		Long:  "get kubeconfig func",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			fmt.Println("get kubeconfig func")

			kubeconfig.RetrieveMgmtKubeconfig()

			return nil
		},
	}
	return c
}
