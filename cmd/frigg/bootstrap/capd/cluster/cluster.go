package cluster

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/clusterapi"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies"
	"github.com/PatrickLaabs/frigg/internal/runtime"
	"github.com/PatrickLaabs/frigg/pkg/common/kubeconfig"
	"github.com/PatrickLaabs/frigg/pkg/common/postbootstrap"
	"github.com/PatrickLaabs/frigg/pkg/common/wait"
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
	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

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
		"templates/kind-config.yaml",
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

	// Get GITHUB_TOKEN environment var
	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Println("Missing Github Token, please set it. Exiting now.")
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_TOKEN")
		fmt.Printf("Github Token:%v\n", os.Getenv("GITHUB_TOKEN"))
	}

	// Create working directory for frigg $HOME/.frigg
	workdir.CreateDir()
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

	// 1. Create a kind cluster, install capi components, clustername is = bootstrapcluster
	wait.Wait(10 * time.Second)
	clusterapi.ClusterAPI()

	// 2. install cni helm chart proxy to the bootstrapcluster
	wait.Wait(10 * time.Second)
	helmchartproxies.InstallBootstrapHelmCharts()

	// 3. generate a manifest, named frigg-mgmt-cluster
	//wait.Wait(10 * time.Second)
	//manifestgenerator.ManifestGeneratorMgmt()

	// 4. modify the manifest of mgmt, to add the helmchart labels to it
	//manifestmodifier.ModifyMgmt()

	// 5. apply the frigg-mgmt-cluster manifest to the bootstrap cluster
	clusterapi.KubectlApplyMgmt()

	// 6. retrieve kubeconfig for the argphub-mgmt-cluster from the bootstrap cluster
	wait.Wait(10 * time.Second)
	kubeconfig.RetrieveMgmtKubeconfig()

	// 8. modify the kubeconfig, to continue working with it
	wait.Wait(5 * time.Second)
	err = kubeconfig.ModifyMgmtKubeconfig()
	if err != nil {
		return err
	}

	// 9. install capi components (like on step 1.) to the frigg-mgmt-cluster
	wait.Wait(5 * time.Second)
	clusterapi.ClusterAPIMgmt()

	// 10. move components from bootstrap cluster to the frigg-mgmt-cluster
	wait.Wait(5 * time.Second)
	clusterapi.Pivot()

	// 11. delete the bootstrap cluster
	postbootstrap.DeleteBootstrapcluster()

	// Installs the HelmChartProxies onto the mgmt-cluster
	wait.Wait(10 * time.Second)
	helmchartproxies.InstallMgmtHelmCharts()

	// 12. generate a workload-cluster manifest
	// => for the sake, we will provide a sample and only apply it onto the mgmt cluster
	// 13. modify the generated manifest with the needed helmchartproxy labels

	// 14. apply the manifest to the frigg-mgmt-cluster
	wait.Wait(5 * time.Second)
	clusterapi.KubectlApplyWorkload()

	// 15. retrieve the kubeconfig
	wait.Wait(10 * time.Second)
	kubeconfig.RetrieveWorkloadKubeconfig()

	// 16. modify the kubeconfig
	wait.Wait(5 * time.Second)
	err = kubeconfig.ModifyWorkloadKubeconfig()
	if err != nil {
		return err
	}

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
