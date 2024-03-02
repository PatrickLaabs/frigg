package clusterapi

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Pivot() {
	println(color.GreenString("Moving clusterapi components from bootstrap to mgmt cluster.."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	bootstrapKubeconfig := homedir + "/" + vars.FriggDirName + "/" + vars.BootstrapkubeconfigName

	mgmtKubeconfig := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	cmd := exec.Command("clusterctl", "--kubeconfig", bootstrapKubeconfig,
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
