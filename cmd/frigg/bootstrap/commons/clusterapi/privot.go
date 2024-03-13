package clusterapi

import (
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

func Pivot() {
	println(color.GreenString("Moving clusterapi components from bootstrap to mgmt cluster.."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	clusterctlPath := filepath.Join(friggToolsDir, "clusterctl")

	bootstrapKubeconfig := filepath.Join(friggDir, vars.BootstrapkubeconfigName)
	mgmtKubeconfig := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(clusterctlPath, "--kubeconfig", bootstrapKubeconfig,
		"move", "--to-kubeconfig", mgmtKubeconfig,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running clusterctl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
