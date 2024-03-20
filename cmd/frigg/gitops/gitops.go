package gitops

import (
	"github.com/PatrickLaabs/frigg/cmd"
	create_template "github.com/PatrickLaabs/frigg/cmd/frigg/gitops/create-template"
	"github.com/PatrickLaabs/frigg/pkg/errors"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for get
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "gitops",
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
	c.AddCommand(create_template.NewCommand(logger, streams))
	return c
}
