package manifestgenerator

import (
	"fmt"
	"os"
	"os/exec"
)

func ManifestGeneratorMgmt() {
	fmt.Println("Generating ClusterAPI Manifest and deploying it to the cluster..")

	homedir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	argohubDirName := ".argohub"
	kubeconfigName := "bootstrapcluster.kubeconfig"
	argohubDir := homedir + "/" + argohubDirName

	// /home/patricklaabs/.argohub/argohub-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName
	fmt.Println("Path to Kubeconfig File:", kubeconfigFlagPath)

	//err = os.Setenv("HELM_VALUES", "")
	//if err != nil {
	//	return
	//}

	//cmd := exec.Command("clusterctl", "generate", "cluster", "workload-cluster",
	//	"--infrastructure", "vcluster:v0.1.3",
	//	"--kubernetes-version", "v1.23.0",
	//	"--kubeconfig", kubeconfigFlagPath,
	//)

	cmd := exec.Command("clusterctl", "generate", "cluster", "argohubmgmtcluster",
		"--flavor", "development",
		"--infrastructure", "docker:v1.6.1",
		"--kubernetes-version", "v1.28.0",
		"--control-plane-machine-count", "3",
		"--worker-machine-count", "3",
		"--kubeconfig", kubeconfigFlagPath,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))

	err = os.WriteFile(argohubDir+"/"+"argohubmgmtclusterManifest.yaml", output, 0750)
	if err != nil {
		return
	}
}
