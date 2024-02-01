package clusterapi

import (
	"fmt"
	"os/exec"
)

//clusterctl generate cluster mgmt-000 --flavor development \
//--kubernetes-version v1.28.6 \
//--control-plane-machine-count=1 \
//--worker-machine-count=3

func ManifestGenerator() {
	fmt.Println("Installing ClusterAPI Components to the KinD Cluster..")

	// Installs CAPI v1.6.1, vCluster v0.1.3 and CAAPH v0.1.1-alpha.0
	cmd := exec.Command("clusterctl", "generate", "cluster", "argohub-cluster",
		"--flavor", "development",
		"--kubernetes-version", "v1.28.6",
		"--control-plane-machine-count", "1",
		"--worker-machine-count", "3",
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
