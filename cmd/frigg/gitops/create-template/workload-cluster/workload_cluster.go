package workload_cluster

import (
	"github.com/PatrickLaabs/frigg/internal/cli"
	"github.com/PatrickLaabs/frigg/internal/reporender"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

type flagpole struct {
	DesiredRepoName string
}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	flags := &flagpole{}
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "workload-cluster",
		Short: "creates gitops template repo for a workload cluster",
		Long:  "creates gitops template repo for a workload cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.OverrideDefaultName(cmd.Flags())
			return runE(flags)
		},
	}
	c.Flags().StringVar(
		&flags.DesiredRepoName,
		"desired-repo-name",
		"",
		"define the name of your generated gitops repo",
	)
	return c
}

func runE(flags *flagpole) error {
	if flags.DesiredRepoName == "" {
		println(color.RedString("Please define your repo name with the option '--desired-repo-name'. Exiting."))
		os.Exit(1)
	}
	println(color.GreenString("Generating GitOps Repo Template for your Workload Cluster"))

	desiredRepoName := flags.DesiredRepoName

	reporender.WorkloadRepo(desiredRepoName)

	println(color.GreenString("Successfully generated your Workload Clusters GitOps repo on your GitHut Account"))
	return nil
}
