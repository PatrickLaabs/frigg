package clusterapi

import (
	"fmt"
	"os"
	"os/exec"
)

func KubectlApplyMgmt() {
	fmt.Println("Applying Manifest to the cluster")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	mgmtcluster := homedir + "/" + argohubDirName + "/" + "argohubmgmtclusterManifest.yaml"

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", mgmtcluster,
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

func KubectlApplyWorkload() {
	fmt.Println("Applying Manifest to the cluster")

	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + friggDirname + "/" + kubeconfigName

	mgmtcluster := homedir + "/" + friggDirname + "/" + "workloadcluster.yaml"

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", mgmtcluster,
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

func retrieveToken() (string, error) {
	// Get GITHUB_TOKEN environment var
	var token string

	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Println("Missing Github Token, please set it. Exiting now.")
		os.Exit(1)
	} else {
		token = os.Getenv("GITHUB_TOKEN")
	}

	return token, nil
}

func ApplyGithubSecretMgmt() {
	token, err := retrieveToken()
	if err != nil {
		fmt.Println("Error retrieving token:", err)
		os.Exit(1)
	}
	fmt.Printf("github token: %v\n", token)

	fmt.Println("Applying Github Secret to the mgmt cluster")

	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"
	fromLiteralString := "--from-literal=token=" + token

	kubeconfigFlagPath := homedir + "/" + friggDirname + "/" + kubeconfigName

	// kubectl -n argocd create secret generic github-token --from-literal=
	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "-n", "argo", "create", "secret", "generic",
		"github-token", fromLiteralString,
	)
	fmt.Printf("github secret gen: %s", cmd)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))
}

func ApplyArgoSecretMgmt() {
	fmt.Println("Applying ArgoCD Login Secret to the mgmt cluster")

	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	kubeconfigFlagPath := homedir + "/" + friggDirname + "/" + kubeconfigName

	// kubectl create secret generic argocd-login --from-literal=password=frigg --from-literal=username=admin -n argo
	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "-n", "argo", "create", "secret", "generic",
		"argocd-login",
		"--from-literal=password=$2a$10$UfHxzEstRBKFAiTH0ZlI8u95SOaRBcXDCxBTBxfmOz14FHC6Vv3de",
		"--from-literal=username=admin",
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

func CreateArgoNSMgmt() {
	fmt.Println("Creating Argo Namespace to the mgmt cluster")

	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"
	kubeconfigName := "argohubmgmtcluster.kubeconfig"

	kubeconfigFlagPath := homedir + "/" + friggDirname + "/" + kubeconfigName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "argo",
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

func CreateArgoNSWorkload() {
	fmt.Println("Creating Argo Namespace to the mgmt cluster")

	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"
	kubeconfigName := "workloadcluster.kubeconfig"

	kubeconfigFlagPath := homedir + "/" + friggDirname + "/" + kubeconfigName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "argo",
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
