package mgmtArgocdApps

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Installation() {
	println(color.GreenString("Applying the Mgmt-ArgoCD-Apps HelmChartProxy to the mgmt-cluster..."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName
	helmchartManifests := homedir + "/" + vars.FriggDirName + "/" + vars.ArgoCDAppsMgmt

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
