package modifykubeconfig

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/common/kubeconfig"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "modifykubeconfig",
		Short: "modify kubeconfic func",
		Long:  "modify kubeconfic func",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("modify kubeconfic func")

			err := kubeconfig.ModifyMgmtKubeconfig()
			if err != nil {
				return err
			}

			return nil
		},
	}
	return c
}
