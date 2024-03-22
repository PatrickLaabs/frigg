package mgmt_cluster

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/internal/cli"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type flagpole struct {
	GitopsTemplateRepo string
}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	flags := &flagpole{}
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "mgmt-cluster",
		Short: "creates gitops template repo for a mgmt cluster",
		Long:  "creates gitops template repo for a mgmt cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.OverrideDefaultName(cmd.Flags())
			return runE(logger, streams, flags)
		},
	}
	c.Flags().StringVar(
		&flags.GitopsTemplateRepo,
		"gitops-template-repo",
		"",
		"add your custom gitops repo template url, like <ORG>/<REPONAME> or <USERNAME>/<REPONAME>",
	)
	return c
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {
	println(color.GreenString("Generating GitOps Repo Template for your Management Cluster"))

	// Steps:
	// Clone the base template repo locally
	//

	println(color.GreenString("Successfully generated your Management Clusters GitOps repo locally at: %s"))
	return nil
}
