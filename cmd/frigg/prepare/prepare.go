package prepare

import (
	"github.com/PatrickLaabs/frigg/cmd/frigg/prepare/download"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "prepare",
		Short: "Prepares everything to use Frigg",
		Long:  "Prepares everything to use Frigg",
		RunE: func(cmd *cobra.Command, args []string) error {
			println(color.GreenString("Downloading Tools you need to operate with Frigg into your home directory."))

			download.GithubCli()
			download.Kubectl()
			download.Clusterctl()
			download.K9s()

			return nil
		},
	}
	return c
}
