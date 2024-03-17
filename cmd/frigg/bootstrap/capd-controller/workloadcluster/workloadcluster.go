package workloadcluster

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/internal/clusterapi"
	"github.com/PatrickLaabs/frigg/internal/generate"
	"github.com/PatrickLaabs/frigg/internal/kubeconfig"
	"github.com/PatrickLaabs/frigg/internal/statuscheck"
	"github.com/PatrickLaabs/frigg/internal/wait"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"time"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "workloadcluster",
		Short: "deploy workload cluster",
		Long:  "deploy workload cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			println(color.GreenString("deploying workload cluster"))

			// Generates a workload-cluster manifest
			// Modifies the manifest of the workload cluster, to add the helmchart labels to it
			wait.Wait(5 * time.Second)
			generate.WorkloadManifest()

			// Applies the workload cluster manifest to the frigg-mgmt-cluster
			wait.Wait(5 * time.Second)
			clusterapi.KubectlApplyWorkload()

			// Retrieves the kubeconfig, like we did for the management cluster previously.
			wait.Wait(10 * time.Second)
			kubeconfig.RetrieveWorkloadKubeconfig()

			// Modifies the kubeconfig, same pattern applies like for the management cluster.
			wait.Wait(5 * time.Second)
			err := kubeconfig.ModifyWorkloadKubeconfig()
			if err != nil {
				fmt.Printf("Error on modifications of the workload cluster kubeconfig: %v\n", err)
			}

			statuscheck.ConditionCniWorkload()

			clusterapi.CreateArgoNSWorkload()

			fmt.Println("Workload Cluster has been successfully provisioned onto your management cluster!")
			return nil
		},
	}
	return c
}
