package clusterapi

import (
	"github.com/PatrickLaabs/frigg/internal/consts"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/fatih/color"
	"os/exec"
	"path/filepath"
)

func ClusterAPI() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(clusterctlPath, "init",
		"--bootstrap", consts.KubeadmVersion,
		"--control-plane", consts.KubeadmVersion,
		"--core", consts.ClusterApiVersion,
		"--infrastructure", consts.DockerInfraVersion,
		"--infrastructure", consts.VClusterInfraVersion,
		"--addon", consts.CaaphVersion,
		"--kubeconfig", kubeconfigFlagPath,
		"--wait-providers",
		"--config", clusterconfigPath,
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
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(clusterctlPath, "init",
		"--bootstrap", consts.KubeadmVersion,
		"--control-plane", consts.KubeadmVersion,
		"--core", consts.ClusterApiVersion,
		"--infrastructure", consts.DockerInfraVersion,
		"--infrastructure", consts.VClusterInfraVersion,
		"--addon", consts.CaaphVersion,
		"--kubeconfig", kubeconfigFlagPath,
		"--wait-providers",
		"--config", clusterconfigPath,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running clusterctl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
