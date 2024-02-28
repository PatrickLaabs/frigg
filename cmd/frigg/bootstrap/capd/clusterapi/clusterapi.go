package clusterapi

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func ClusterAPI() {
	println(color.GreenString("Installing ClusterAPI Components to the KinD Cluster.."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + bootstrapkubeconfigName

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
		println(color.RedString("Error running clusterctl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ClusterAPIMgmt() {
	println(color.GreenString("Installing ClusterAPI Components to the Management Cluster.."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + managementKubeconfigName

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
		println(color.RedString("Error running clusterctl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
