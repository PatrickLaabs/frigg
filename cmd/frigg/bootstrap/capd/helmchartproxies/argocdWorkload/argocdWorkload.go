package argocdWorkload

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

func Installation() {
	println(color.GreenString("Applying the Helmchartproxy for ArgoCD for Workload Clusters to the mgmt-cluster..."))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, "kubectl")

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)
	helmchartManifests := filepath.Join(friggDir, vars.ArgoCDWorkload)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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
