package mgmtArgocdWorkflows

import (
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Installation() {
	println(color.GreenString("Applying the Mgmt-ArgoCD-Workflows HelmChartProxy to the mgmt-cluster..."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDirName := helmchartproxies.FriggDirName
	managementKubeconfigName := helmchartproxies.ManagementKubeconfigName

	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + managementKubeconfigName

	helmchartManifests := "templates/helmchartproxies/mgmt-argocd-workflows.yaml"

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
