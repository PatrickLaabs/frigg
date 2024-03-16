package clusterapi

import (
	"github.com/PatrickLaabs/frigg/pkg/consts"
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

var clusterctl = "clusterctl_" + consts.ClusterctlVersion

func ClusterAPI() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	clusterctlPath := filepath.Join(friggToolsDir, clusterctl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)
	clusterconfigPath := filepath.Join(friggDir, vars.ClusterctlConfigName)

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
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	clusterctlPath := filepath.Join(friggToolsDir, clusterctl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)
	clusterconfigPath := filepath.Join(friggDir, vars.ClusterctlConfigName)

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
