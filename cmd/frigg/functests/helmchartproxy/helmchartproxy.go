package helmchartproxy

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/tmpl/helmchartsproxies"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "hc",
		Short: "modify hc func",
		Long:  "modify hc func",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("modify hc func")

			helmchartsproxies.Cni()

			return nil
		},
	}
	return c
}
