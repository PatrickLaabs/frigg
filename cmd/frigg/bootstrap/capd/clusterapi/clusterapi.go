package clusterapi

import (
	"fmt"
	"os"
	"os/exec"
)

func ClusterAPI() {
	fmt.Println("Installing ClusterAPI Components to the KinD Cluster..")

	homedir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	// Installs CAPI v1.6.1, Docker and CAAPH v0.1.1-alpha.0
	cmd := exec.Command("clusterctl", "init",
		"--bootstrap", "kubeadm:v1.6.1",
		"--control-plane", "kubeadm:v1.6.1",
		"--core", "cluster-api:v1.6.1",
		"--infrastructure", "docker:v1.6.1",
		"--infrastructure", "vcluster:v0.1.3",
		"--addon", "helm:v0.1.1-alpha.0",
		"--kubeconfig", kubeconfigFlagPath,
		"--wait-providers",
		"--config", "templates/clusterctl-config.yaml",
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

func ClusterAPIMgmt() {
	fmt.Println("Installing ClusterAPI Components to the KinD Cluster..")

	homedir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	argohubDirName := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	// Installs CAPI v1.6.1, Docker and CAAPH v0.1.1-alpha.0
	cmd := exec.Command("clusterctl", "init",
		"--bootstrap", "kubeadm:v1.6.1",
		"--control-plane", "kubeadm:v1.6.1",
		"--core", "cluster-api:v1.6.1",
		"--infrastructure", "docker:v1.6.1",
		"--infrastructure", "vcluster:v0.1.3",
		"--addon", "helm:v0.1.1-alpha.0",
		"--kubeconfig", kubeconfigFlagPath,
		"--wait-providers",
		"--config", "templates/clusterctl-config.yaml",
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
