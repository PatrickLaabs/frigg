package mgmtArgocdApps

import (
	"fmt"
	"os"
	"os/exec"
)

func Installation() {
	fmt.Println("Applying HelmChartProxies to the mgmt-cluster...")

	homedir, _ := os.UserHomeDir()

	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + kubeconfigName

	//helmchartManifests := "templates/helmchartproxies/mgmt-argocd-apps.yaml"
	newFile := "mgmt-argocd-apps.yaml"
	helmchartManifests := friggDir + "/" + newFile

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
