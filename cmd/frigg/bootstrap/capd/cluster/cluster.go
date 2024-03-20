package cluster

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/internal/clusterapi"
	"github.com/PatrickLaabs/frigg/internal/directories"
	"github.com/PatrickLaabs/frigg/internal/generate"
	"github.com/PatrickLaabs/frigg/internal/helmchart"
	"github.com/PatrickLaabs/frigg/internal/kubeconfig"
	"github.com/PatrickLaabs/frigg/internal/postbootstrap"
	"github.com/PatrickLaabs/frigg/internal/prepare"
	"github.com/PatrickLaabs/frigg/internal/reporender"
	"github.com/PatrickLaabs/frigg/internal/runtime"
	"github.com/PatrickLaabs/frigg/internal/sshkey"
	"github.com/PatrickLaabs/frigg/internal/statuscheck"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/PatrickLaabs/frigg/internal/wait"
	"github.com/fatih/color"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/pkg/cluster"
	"github.com/PatrickLaabs/frigg/pkg/errors"
	"github.com/PatrickLaabs/frigg/pkg/log"

	"github.com/PatrickLaabs/frigg/internal/cli"
)

type flagpole struct {
	Name                   string
	Config                 string
	ImageName              string
	Retain                 bool
	Wait                   time.Duration
	Kubeconfig             string
	GitopsTemplateRepo     string
	GitopsWorkloadTemplate string
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
		Use:   "cluster",
		Short: "Creates a local Kubernetes cluster",
		Long:  "Creates a local Kubernetes cluster using Docker container 'nodes'",
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
	c.Flags().StringVar(
		&flags.GitopsWorkloadTemplate,
		"gitops-workload-template-repo",
		"",
		"add your custom gitops repo template url, like <ORG>/<REPONAME> or <USERNAME>/<REPONAME>",
	)
	return c
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {

	if os.Getenv("GITHUB_TOKEN") == "" {
		println(color.RedString("Missing Github Token, please set it. Exiting now."))
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_TOKEN")
	}

	if os.Getenv("GITHUB_USERNAME") == "" {
		println(color.RedString("Missing Github Username, please set it. Exiting now."))
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_USERNAME")
	}

	if os.Getenv("GITHUB_MAIL") == "" {
		println(color.RedString("Missing Github User Mail, please set it. Exiting now."))
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_MAIL")
	}

	if flags.GitopsTemplateRepo == "" {
		println(color.YellowString("No Gitops Template Repo specified, using default: %s ", vars.FriggMgmtTemplateName))
		flags.GitopsTemplateRepo = vars.FriggMgmtTemplateName
	}

	if flags.GitopsWorkloadTemplate == "" {
		println(color.YellowString("No Gitops Template Repo specified, using default: %s ", vars.FriggWorkloadTemplateName))
		flags.GitopsWorkloadTemplate = vars.FriggWorkloadTemplateName
	}

	// Creating Working, tools and controllers directories
	directories.Create()

	// Preparing the CLIs for Frigg
	prepare.Tools()

	// Generating kind-config
	generate.KindConfigGen()

	// Generating SSH Key pair
	sshkey.KeypairGen()

	// Generating clusterctl config
	generate.ClusterctlConfig()

	// Generating HelmChartProxies
	println(color.GreenString("Generating Helm Chart Proxy Manifests.."))
	generate.Cni()
	generate.Vault()
	generate.ArgoCDWorkloadClusters()
	generate.ArgoWorkflows()
	generate.ArgoRollouts()
	generate.ArgoEvents()
	generate.MgmtArgoCD()
	generate.MgmtArgoApps()

	// Generates a manifest for the management cluster, named frigg-mgmt-cluster
	generate.MgmtManifest()

	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
		runtime.GetDefault(logger),
	)

	// handle config flag, we might need to read from stdin
	withConfig, err := configOption(flags.Config, streams.In)
	if err != nil {
		return err
	}

	// create the cluster
	if err = provider.Create(
		flags.Name,
		withConfig,
		cluster.CreateWithNodeImage(flags.ImageName),
		cluster.CreateWithRetain(flags.Retain),
		cluster.CreateWithWaitForReady(flags.Wait),
		cluster.CreateWithKubeconfigPath(flags.Kubeconfig),
		cluster.CreateWithDisplayUsage(true),
		cluster.CreateWithDisplaySalutation(true),
	); err != nil {
		return errors.Wrap(err, "failed to create cluster")
	}

	// Rendering gitops repo
	reporender.FullStage(flags.GitopsTemplateRepo, flags.GitopsWorkloadTemplate)

	// Installs capi components on the bootstrap cluster.
	println(color.GreenString("Applying ClusterAPI and Cert-Manager Components and waiting for the 'Ready' condition.."))
	clusterapi.ClusterAPI()
	statuscheck.ConditionsCertManagers()
	statuscheck.ConditionsCapiControllers()
	statuscheck.ConditionsCapdControllers()
	statuscheck.ConditionsCaaphControllers()

	// Installs a CNI solution helm chart proxy to the bootstrapcluster
	// This is needed, to make the worker nodes ready and complete the bootstrap deployment
	helmchart.InstallOnBootstrap()

	// Applies the frigg-mgmt-cluster manifest to the bootstrap cluster
	// to create the first 'real' management cluster
	clusterapi.KubectlApplyMgmt()

	// Retrieves the kubeconfig for the frigg-mgmt-cluster from the bootstrap cluster
	// so that we can later on use the kubeconfig to target the correct cluster for deployments.
	wait.Wait(2 * time.Second)
	kubeconfig.RetrieveMgmtKubeconfig()

	// Modifes the kubeconfig, to let us interact with the newly created kubernetes cluster.
	// On MacOS, there is an issue, where you need to replace ip and the port address, in order to
	// successfully connect to the cluster.
	// ToDo:
	// We shall check, if the user is running on macOS, Linux and/or Windows.
	// Depending on the OS, the modification, of the kubeconfig, may not be needed.
	err = kubeconfig.ModifyMgmtKubeconfig()
	if err != nil {
		fmt.Printf("Error on modification of mgmt clusters kubeconfig: %v\n", err)
	}

	println(color.GreenString("Waiting for 'Ready' conditions for Tigera, CoreDNS and CNI.."))
	statuscheck.ConditionTigerOperatorMgmt()
	statuscheck.ConditionCoreDnsMgmt()
	statuscheck.ConditionsCni()

	// Creates the argo namespace on the Mgmt Cluster
	clusterapi.CreateArgoNSMgmt()

	// Installs the capi components to the frigg-mgmt-cluster
	// This part may take a while.
	println(color.GreenString("Installation of ClusterAPI Componentes and waiting for the 'Ready' condition.."))
	clusterapi.ClusterAPIMgmt()
	statuscheck.ConditionsCertManagersMgmt()
	statuscheck.ConditionsCapiControllersMgmt()
	statuscheck.ConditionsCapdControllersMgmt()
	statuscheck.ConditionsCaaphControllersMgmt()

	// Github Token Secret deployment
	clusterapi.ApplyGithubSecretMgmt()
	// ArgoCD Default Login Secret deployment
	clusterapi.ApplyArgoSecretMgmt()

	// Installs the HelmChartProxies onto the mgmt-cluster
	helmchart.InstallOnMgmt()

	// Moves the capi components from the bootstrap cluster to the frigg-mgmt-cluster
	wait.Wait(15 * time.Second)
	clusterapi.Pivot()

	// Deletes the bootstrap cluster, since we don't need it any longer
	// and to free up some hardware resources.
	postbootstrap.DeleteBootstrapcluster()
	println(color.GreenString("Successfully provisioned your management cluster ontop of commons."))
	return nil
}

// configOption converts the raw --config flag value to a cluster creation
// option matching it. it will read from stdin if the flag value is `-`
func configOption(rawConfigFlag string, stdin io.Reader) (cluster.CreateOption, error) {
	// if not - then we are using a real file
	if rawConfigFlag != "-" {
		return cluster.CreateWithConfigFile(rawConfigFlag), nil
	}
	// otherwise read from stdin
	raw, err := io.ReadAll(stdin)
	if err != nil {
		return nil, errors.Wrap(err, "error reading config from stdin")
	}
	return cluster.CreateWithRawConfig(raw), nil
}
