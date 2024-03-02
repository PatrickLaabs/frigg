package clusterapi

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func KubectlApplyMgmt() {
	println(color.GreenString("Applying Manifest to the cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.BootstrapkubeconfigName

	mgmtcluster := homedir + "/" + vars.FriggDirName + "/" + vars.MgmtManifest

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", mgmtcluster,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func KubectlApplyWorkload() {
	fmt.Println("Applying Manifest to the cluster")
	println(color.GreenString("Applying workload cluster manifest to the management cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	workloadcluster := homedir + "/" + vars.FriggDirName + "/" + vars.WorkloadManifest

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "apply",
		"-f", workloadcluster,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func retrieveToken() (string, error) {
	var token string

	if os.Getenv("GITHUB_TOKEN") == "" {
		println(color.RedString("Missing Github Token, please set it. Exiting now."))
		os.Exit(1)
	} else {
		token = os.Getenv("GITHUB_TOKEN")
	}

	return token, nil
}

func ApplyGithubSecretMgmt() {
	println(color.GreenString("Applying Github Secret to the mgmt cluster"))

	token, err := retrieveToken()
	if err != nil {
		println(color.RedString("Error retrieving token: %v\n", err))
		os.Exit(1)
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	fromLiteralString := "--from-literal=token=" + token
	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "-n", "argo", "create", "secret", "generic",
		"github-token", fromLiteralString,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyArgoSecretMgmt() {
	println(color.GreenString("Applying ArgoCD Login Secret to the mgmt cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "-n", "argo", "create", "secret", "generic",
		"argocd-login",
		"--from-literal=password=$2a$10$UfHxzEstRBKFAiTH0ZlI8u95SOaRBcXDCxBTBxfmOz14FHC6Vv3de",
		"--from-literal=username=admin",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateArgoNSMgmt() {
	println(color.GreenString("Creating Argo Namespace to the mgmt cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.ManagementKubeconfigName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "argo",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateArgoNSWorkload() {
	println(color.GreenString("Creating Argo Namespace to the mgmt cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	kubeconfigFlagPath := homedir + "/" + vars.FriggDirName + "/" + vars.WorkloadKubeconfigName

	cmd := exec.Command("kubectl", "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "argo",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
