package cni

import (
	"fmt"
	"os"
	"os/exec"
)

func Installation() {
	fmt.Println("Applying CNI HelmChartProxies to the mgmt-cluster...")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

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
