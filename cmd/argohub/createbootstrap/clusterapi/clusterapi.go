package clusterapi

import (
	"fmt"
	"os/exec"
)

func ClusterAPI() {
	fmt.Println("Installing ClusterAPI Components to the KinD Cluster..")

	// Installs CAPI v1.6.1, Docker and CAAPH v0.1.1-alpha.0
	cmd := exec.Command("clusterctl", "init",
		"--core", "cluster-api:v1.6.1",
		"--bootstrap", "kubeadm:v1.6.1",
		"--control-plane", "kubeadm:v1.6.1",
		"--infrastructure", "docker",
		"--addon", "helm:v0.1.1-alpha.0",
		"--kubeconfig", "test",
	)

	// Set the environment variable
	cmd.Env = append(cmd.Env, "CLUSTER_TOPOLOGY=true")
	cmd.Env = append(cmd.Env, "EXP_MACHINE_POOL=true")

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))

	//VclusterProvider()
	//ManifestGenerator()
}
