package clusterapi

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/consts"
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

var kubectl = "kubectl_" + consts.KubectlVersion

func KubectlApplyMgmt() {
	println(color.GreenString("Applying Manifest to the cluster"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)
	mgmtcluster := filepath.Join(friggDir, vars.MgmtManifest)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)
	workloadcluster := filepath.Join(friggDir, vars.WorkloadManifest)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	fromLiteralString := "--from-literal=token=" + token
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	println(kubectlPath)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.WorkloadKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

func CreateCapiNs() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capi-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateCapdNs() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capd-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateCaaphNs() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "caaph-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateKubeadmBootstrapNs() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capi-kubeadm-bootstrap-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateKubeAdmControlPlaneNs() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capi-kubeadm-control-plane-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyCoreProvider() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	path := filepath.Join(friggControllerDir, vars.CoreProviderName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyControlPlaneProv() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	path := filepath.Join(friggControllerDir, vars.ControlPlaneProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyBootstrapProv() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	path := filepath.Join(friggControllerDir, vars.BootstrapProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyDockerInfraProv() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	path := filepath.Join(friggControllerDir, vars.DockerInfraProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyAddonHelmProv() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	path := filepath.Join(friggControllerDir, vars.HelmAddonProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

// ApplyCertManager applies cert-manager installation on the Bootstrap Cluster
func ApplyCertManager() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", vars.CertManagerManifest,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

// ApplyCapiOperator applies the ClusterAPI Operator to the bootstrap cluster
func ApplyCapiOperator() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", vars.CapiOperatorManifest,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateCapiNsMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capi-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateCapdNsMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capd-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateCaaphNsMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "caaph-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateKubeadmBootstrapNsMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capi-kubeadm-bootstrap-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateKubeAdmControlPlaneNsMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)

	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "capi-kubeadm-control-plane-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyCoreProviderMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	path := filepath.Join(friggControllerDir, vars.CoreProviderName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyControlPlaneProvMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	path := filepath.Join(friggControllerDir, vars.ControlPlaneProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyBootstrapProvMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	path := filepath.Join(friggControllerDir, vars.BootstrapProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyDockerInfraProvMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	path := filepath.Join(friggControllerDir, vars.DockerInfraProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func ApplyAddonHelmProvMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	friggControllerDir := filepath.Join(friggDir, vars.ControllerDir)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	path := filepath.Join(friggControllerDir, vars.HelmAddonProvName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", path,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

// ApplyCertManagerMgmt applies cert-manager installation on the mgmt Cluster
func ApplyCertManagerMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", vars.CertManagerManifest,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

// ApplyCapiOperatorMgmt applies the ClusterAPI Operator on the mgmt cluster
func ApplyCapiOperatorMgmt() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir := filepath.Join(friggDir, vars.FriggTools)
	kubectlPath := filepath.Join(friggToolsDir, kubectl)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "apply", "-f", vars.CapiOperatorManifest,
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
