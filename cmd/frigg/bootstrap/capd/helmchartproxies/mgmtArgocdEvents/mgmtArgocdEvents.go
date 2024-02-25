package mgmtArgocdEvents

import (
	"fmt"
	"os"
	"os/exec"
)

func Installation() {
	fmt.Println("Applying HelmChartProxies to the mgmt-cluster...")

	homedir, _ := os.UserHomeDir()

	friggDirName := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + kubeconfigName

	helmchartManifests := "templates/helmchartproxies/mgmt-argocd-events.yaml"

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", helmchartManifests,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))
}
