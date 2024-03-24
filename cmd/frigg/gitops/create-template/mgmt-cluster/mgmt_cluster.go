package mgmt_cluster

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/internal/cli"
	"github.com/PatrickLaabs/frigg/internal/consts"
	"github.com/PatrickLaabs/frigg/internal/reporender"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type flagpole struct {
	DesiredRepoName string
}

var (
	gh                     = "gh_" + consts.GithubCliVersion
	homedir, _             = os.UserHomeDir()
	friggDir               = filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir          = filepath.Join(friggDir, vars.FriggTools)
	ghCliPath              = filepath.Join(friggToolsDir, gh)
	sshpublickeyPath       = filepath.Join(friggDir, vars.PublickeyName)
	localRepo              = filepath.Join(friggDir, vars.RepoName)
	localRepoStoragePath   = filepath.Join(friggDir, vars.RepoName)
	gitopsWorkloadTemplate = vars.FriggWorkloadTemplateName
)

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
		&flags.DesiredRepoName,
		"desired-repo-name",
		"",
		"define the name of your generated gitops repo",
	)
	return c
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {
	if flags.DesiredRepoName == "" {
		println(color.RedString("Please define your repo name with the option '--desired-repo-name'. Exiting."))
		os.Exit(1)
	}
	println(color.GreenString("Generating GitOps Repo Template for your Management Cluster"))

	desiredRepoName := flags.DesiredRepoName

	reporender.MgmtRepo(desiredRepoName)

	println(color.GreenString("Successfully generated your Management Clusters GitOps repo on your GitHut Account"))
	return nil
}
