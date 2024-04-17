package clusterapi

import (
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

func KubectlApplyMgmt() {
	println(color.GreenString("Applying Manifest to the cluster"))

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
	println(color.GreenString("Applying workload cluster manifest to the management cluster"))

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

func CreateVclusterProvNs() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "cluster-api-provider-vcluster-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateVclusterNs() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "vcluster",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateVclusterProvNsMgmt() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "cluster-api-provider-vcluster-system",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error running kubectl: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}

func CreateVclusterNsMgmt() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
		kubeconfigFlagPath, "create", "namespace", "vcluster",
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

func VclusterInfraProv() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	path := filepath.Join(friggControllerDir, vars.VclusterProvName)

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

func VclusterInfraProvMgmt() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	path := filepath.Join(friggControllerDir, vars.VclusterProvName)

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
