package mgmt_cluster

import (
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/internal/cli"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

type flagpole struct {
	Name               string
	Config             string
	ImageName          string
	Retain             bool
	Wait               time.Duration
	Kubeconfig         string
	GitopsTemplateRepo string
}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing users home directory: %v\n", err))
	}

	kubeconfigFlagPath := filepath.Join(homedir, vars.FriggDirName, vars.BootstrapkubeconfigName)
	kindconfigPath := filepath.Join(homedir, vars.FriggDirName, vars.KindconfigName)
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
	c.Flags().StringVarP(
		&flags.Name,
		"name",
		"n",
		"bootstrapcluster",
		"cluster name, overrides KIND_CLUSTER_NAME, config (default kind)",
	)
	c.Flags().StringVar(
		&flags.Config,
		"config",
		kindconfigPath,
		"path to a kind config file",
	)
	c.Flags().StringVar(
		&flags.ImageName,
		"image",
		"",
		"node docker image to use for booting the cluster",
	)
	c.Flags().BoolVar(
		&flags.Retain,
		"retain",
		false,
		"retain nodes for debugging when cluster creation fails",
	)
	c.Flags().DurationVar(
		&flags.Wait,
		"wait",
		time.Duration(0),
		"wait for control plane node to be ready (default 0s)",
	)
	c.Flags().StringVar(
		&flags.Kubeconfig,
		"kubeconfig",
		kubeconfigFlagPath,
		"sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config",
	)
	c.Flags().StringVar(
		&flags.GitopsTemplateRepo,
		"gitops-template-repo",
		"",
		"add your custom gitops repo template url, like <ORG>/<REPONAME> or <USERNAME>/<REPONAME>",
	)
	return c
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {
	return nil
}
