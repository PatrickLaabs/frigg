package clusterapi

import (
	"fmt"
	"os"
	"os/exec"
)

func Pivot() {
	fmt.Println("Moving clusterapi components from bootstrap to mgmt cluster..")
	homedir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"
	targetKubeconfigName := "argohubmgmtcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	bootstrapKubeconfig := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Printf("bootstrap kubeconfig path: %s\n", bootstrapKubeconfig)

	mgmtKubeconfig := homedir + "/" + argohubDirName + "/" + targetKubeconfigName
	fmt.Printf("mgmt cluster kubeconfig path: %s\n", mgmtKubeconfig)

	// clusterctl --kubeconfig bootstrapcluster.kubeconfig move --to-kubeconfig argohubmgmtcluster.kubeconfig
	cmd := exec.Command("clusterctl", "--kubeconfig", bootstrapKubeconfig,
		"move", "--to-kubeconfig", mgmtKubeconfig,
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
