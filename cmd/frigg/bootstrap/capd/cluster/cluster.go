package cluster

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/clusterapi"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/argocdWorkload"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/cni"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/cnibootstrap"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdApps"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdEvents"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdRollouts"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgocdWorkflows"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtArgohub"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies/mgmtVault"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/reporender"
	"github.com/PatrickLaabs/frigg/internal/runtime"
	"github.com/PatrickLaabs/frigg/pkg/common/kubeconfig"
	"github.com/PatrickLaabs/frigg/pkg/common/postbootstrap"
	"github.com/PatrickLaabs/frigg/pkg/common/sshkey"
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/PatrickLaabs/frigg/pkg/common/wait"
	"github.com/PatrickLaabs/frigg/tmpl/clusterctlconfig"
	"github.com/PatrickLaabs/frigg/tmpl/helmchartsproxies"
	"github.com/PatrickLaabs/frigg/tmpl/kindconfig"
	"github.com/PatrickLaabs/frigg/tmpl/mgmtmanifestgen"
	"github.com/fatih/color"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/pkg/cluster"
	"github.com/PatrickLaabs/frigg/pkg/common/workdir"
	"github.com/PatrickLaabs/frigg/pkg/errors"
	"github.com/PatrickLaabs/frigg/pkg/log"

	"github.com/PatrickLaabs/frigg/internal/cli"
)

type flagpole struct {
	Name       string
	Config     string
	ImageName  string
	Retain     bool
	Wait       time.Duration
	Kubeconfig string
}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing users home directory: %v\n", err))
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.BootstrapkubeconfigName
	kindconfigPath := homedir + "/" + vars.FriggDirName + "/" + vars.KindconfigName
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
	return c
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {

	if os.Getenv("GITHUB_TOKEN") == "" {
		println(color.RedString("Missing Github Token, please set it. Exiting now."))
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_TOKEN")
		println(color.GreenString("Found Github Token Environment variable. Continuing.."))
	}

	if os.Getenv("GITHUB_USERNAME") == "" {
		println(color.RedString("Missing Github Username, please set it. Exiting now."))
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_USERNAME")
		println(color.GreenString("Found Github Username Environment variable. Continuing.."))
	}

	if os.Getenv("GITHUB_MAIL") == "" {
		println(color.RedString("Missing Github Username, please set it. Exiting now."))
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_MAIL")
		println(color.GreenString("Found Github Username Environment variable. Continuing.."))
	}

	// Create working directory named .frigg inside the users homedirectory.
	workdir.CreateDir()

	// Generating kind-config
	kindconfig.KindConfigGen()

	// Generating SSH Key pair
	sshkey.KeypairGen()

	// Generating clusterctl config
	clusterctlconfig.ClusterctlConfigGen()

	// Generating HelmChartProxies
	helmchartsproxies.Cni()
	helmchartsproxies.Vault()
	helmchartsproxies.ArgoCDWorkloadClusters()
	helmchartsproxies.ArgoWorkflows()
	helmchartsproxies.ArgoRollouts()
	helmchartsproxies.ArgoEvents()
	helmchartsproxies.MgmtArgoCD()
	helmchartsproxies.MgmtArgoApps()

	// Generates a manifest for the management cluster, named frigg-mgmt-cluster
	wait.Wait(10 * time.Second)
	mgmtmanifestgen.Gen()

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
	reporender.FullStage()

	// Installs capi components on the bootstrap cluster.
	// clustername is bootstrapcluster
	wait.Wait(10 * time.Second)
	clusterapi.ClusterAPI()

	// Installs a CNI solution helm chart proxy to the bootstrapcluster
	// This is needed, to make the worker nodes ready and complete the bootstrap deployment
	wait.Wait(10 * time.Second)
	cnibootstrap.Installation()

	// Applies the frigg-mgmt-cluster manifest to the bootstrap cluster
	// to create the first 'real' management cluster
	wait.Wait(5 * time.Second)
	clusterapi.KubectlApplyMgmt()

	// Retrieves the kubeconfig for the frigg-mgmt-cluster from the bootstrap cluster
	// so that we can later on use the kubeconfig to target the correct cluster for deployments.
	wait.Wait(10 * time.Second)
	kubeconfig.RetrieveMgmtKubeconfig()

	// Modifes the kubeconfig, to let us interact with the newly created kubernetes cluster.
	// On MacOS, there is an issue, where you need to replace ip and the port address, in order to
	// successfully connect to the cluster.
	// ToDo:
	// We shall check, if the user is running on macOS, Linux and/or Windows.
	// Depending on the OS, the modification, of the kubeconfig, may not be needed.
	wait.Wait(5 * time.Second)
	err = kubeconfig.ModifyMgmtKubeconfig()
	if err != nil {
		fmt.Printf("Error on modification of mgmt clusters kubeconfig: %v\n", err)
	}

	// Installs the capi components to the frigg-mgmt-cluster
	// This part may take a while.
	wait.Wait(15 * time.Second)
	clusterapi.ClusterAPIMgmt()

	wait.Wait(30 * time.Second)
	// Creates the argo namespace on the Mgmt Cluster
	clusterapi.CreateArgoNSMgmt()

	// Applies the Github Token and the default ArgoCD Login Credentials as a
	// kubernetes secret on the argo namespace.
	// This is needed to let us interact with github, to clone, refactor and push the needed
	// gitops repositories.
	//
	// On the deployment, of new workload kubernetes clusters - which will be attached to the management
	// cluster - we run ArgoCD Workflows, which will create a pod, which runs a script.
	// This script logs in to the argocd instance, and adds the new kubernetes cluster to it, and
	// also adds a label to the cluster, with which we can proceed the automation steps.

	wait.Wait(5 * time.Second)
	// Github Token Secret deployment
	clusterapi.ApplyGithubSecretMgmt()
	// ArgoCD Default Login Secret deployment
	clusterapi.ApplyArgoSecretMgmt()

	// Moves the capi components from the bootstrap cluster to the frigg-mgmt-cluster
	wait.Wait(5 * time.Second)
	clusterapi.Pivot()

	// Deletes the bootstrap cluster, since we don't need it any longer
	// and to free up some hardware resources.
	postbootstrap.DeleteBootstrapcluster()

	// Installs the HelmChartProxies onto the mgmt-cluster
	argocdWorkload.Installation()
	cni.Installation()
	mgmtArgocdApps.Installation()
	mgmtArgocdEvents.Installation()
	mgmtArgocdRollouts.Installation()
	mgmtArgocdWorkflows.Installation()
	mgmtArgohub.Installation()
	mgmtVault.Installation()

	println(color.GreenString("Successfully provisioned your management cluster ontop of capd."))

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
