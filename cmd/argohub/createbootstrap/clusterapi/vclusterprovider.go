package clusterapi

import (
	"fmt"
	"os/exec"
)

func VclusterProvider() {
	fmt.Println("Installing ClusterAPI vCluster Provider to the KinD Cluster..")

	// Installs CAPI v1.6.1, vCluster v0.1.3 and CAAPH v0.1.1-alpha.0
	cmd := exec.Command("clusterctl", "init",
		"--infrastructure", "vcluster:v0.1.3",
		"--kubeconfig", "test",
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
