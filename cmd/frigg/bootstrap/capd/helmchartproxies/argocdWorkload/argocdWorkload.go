package argocdWorkload

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Installation() {
	println(color.GreenString("Applying the Helmchartproxy for ArgoCD for Workload Clusters to the mgmt-cluster..."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	helmchartManifests := homedir + "/" + vars.FriggDirName + "/" + vars.ArgoCDWorkload

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", helmchartManifests,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl apply for argocd workload helmchartproxy: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
