package cluster

import (
	"fmt"
	capi "github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd/clusterapi"
	hc "github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/bootstrap/capd/helmchartproxies"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/internal/runtime"
	w "github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/common/wait"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cluster"
	d "github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/common/workdir"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/errors"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"

	"github.com/PatrickLaabs/cli_clusterapi-argohub/internal/cli"
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

	argohubDirName := ".argohub"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.argohub/argohub-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

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
	d.CreateDir()
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
	w.Wait(10 * time.Second)
	capi.ClusterAPI()

	// 2. install cni helm chart proxy to the bootstrapcluster
	w.Wait(5 * time.Second)
	hc.InstallBootstrapHelmCharts()

	// 3. generate a manifest, named argohub-mgmt-cluster
	//w.Wait(10 * time.Second)
	//manifestgenerator.ManifestGeneratorMgmt()

	// 4. modify the manifest of mgmt, to add the helmchart labels to it
	//manifestmodifier.ModifyMgmt()

	// 5. apply the argohub-mgmt-cluster manifest to the bootstrap cluster
	capi.KubectlApplyMgmt()

	// 6. retrieve kubeconfig for the argphub-mgmt-cluster from the bootstrap cluster
	//	  => clusterctl --kubeconfig bootstrapcluster.kubeconfig get kubeconfig argohubmgmtcluster > argohubmgmtcluster.kubeconfig
	// 7. store the kubeconfig of the argohub-mgmt-cluster to the .argohub folder
	// 8. modify the kubeconfig, to continue working with it
	//    => sed -i -e "s/server:.*/server: https:\/\/$(docker port argohubmgmtcluster-lb 6443/tcp | sed "s/0.0.0.0/127.0.0.1/")/g" ./argohubmgmtcluster.kubebeconfig
	// 9. install capi components (like on step 1.) to the argohub-mgmt-cluster
	// 10. move components from bootstrap cluster to the argohub-mgmt-cluster
	// 11. delete the bootstrap cluster
	// 12. generate a workload-cluster manifest
	// 13. modify the generated manifest with the needed helmchartproxy labels
	// 14. apply the manifest to the argohub-mgmt-cluster
	// 15. retrieve the kubeconfig
	// 16. modify the kubeconfig

	// Installs the HelmChartProxies onto the mgmt-cluster
	//w.Wait(10 * time.Second)
	//hc.InstallMgmtHelmCharts()

	//w.Wait(5 * time.Second)
	//manifestmodifier.Modifyworkload()

	// Applies the prev. generated Manifest of the workload cluster
	//w.Wait(10 * time.Second)
	//capi.KubectlApplyWorkload()

	// ToDo:
	// Retrieve kubeconfig of workload cluster and write it to the .argohub folder

	// ToDo:
	// Modify the retrieved and stored kubeconfig, so it works with CAPD

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
