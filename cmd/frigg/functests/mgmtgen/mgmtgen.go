package mgmtgen

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/PatrickLaabs/frigg/tmpl/mgmtmanifestgen"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "mgmtgen",
		Short: "mgmtgen",
		Long:  "mgmtgen",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("mgmtgen")

			mgmtmanifestgen.Gen()

			return nil
		},
	}
	return c
}
