package clusterapi

import (
	"fmt"
	"os"
	"os/exec"
)

func KubectlApplyMgmt() {
	fmt.Println("Applying Manifest to the cluster")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

	mgmtcluster := homedir + "/" + argohubDirName + "/" + "argohubmgmtclusterManifest.yaml"

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", mgmtcluster,
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

func KubectlApplyWorkload() {
	fmt.Println("Applying Manifest to the cluster")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

	mgmtcluster := homedir + "/" + argohubDirName + "/" + "workloadcluster.yaml"

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", mgmtcluster,
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
