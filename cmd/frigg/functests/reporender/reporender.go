package reporender

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/reporender"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "reporender",
		Short: "reporender",
		Long:  "reporender",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("reporender")

			reporender.FullStage()

			return nil
		},
	}
	return c
}
