package cnibootstrap

import (
	"fmt"
	"os"
	"os/exec"
)

func Installation() {
	fmt.Println("Applying HelmChartProxies to the mgmt-cluster...")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".argohub"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.argohub/argohub-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

	//workloadClusterManifestPath := homedir + "/" + argohubDirName + "/" + "gened-Manifest.yaml"
	helmchartManifests := "templates/helmchartproxies/cni.yaml"

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", helmchartManifests,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))
}
