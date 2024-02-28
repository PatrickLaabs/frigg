package argocdWorkload

import (
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies"
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

	friggDirName := helmchartproxies.FriggDirName
	managementKubeconfigName := helmchartproxies.ManagementKubeconfigName

	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + managementKubeconfigName

	helmchartManifests := "templates/helmchartproxies/argocd-workload-clusters.yaml"

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
