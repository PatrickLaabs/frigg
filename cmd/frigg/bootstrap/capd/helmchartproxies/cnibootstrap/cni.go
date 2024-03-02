package cnibootstrap

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Installation() {
	println(color.GreenString("Applying the CNI Bootstrap HelmChartProxy to the bootstrap-cluster..."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.BootstrapkubeconfigName

	helmchartManifests := homedir + "/" + vars.FriggDirName + "/" + vars.CniName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", helmchartManifests,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
