package helmchart

import (
	"github.com/PatrickLaabs/frigg/internal/consts"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	homedir, _    = os.UserHomeDir()
	friggDir      = filepath.Join(homedir, vars.FriggDirName)
	friggToolsDir = filepath.Join(friggDir, vars.FriggTools)
	kubectl       = "kubectl_" + consts.KubectlVersion
	kubectlPath   = filepath.Join(friggToolsDir, kubectl)
)

// argoCdTargetCluster installs the ArgoCD HelmChartProxy to the mgmt cluster used by the workload clusters
func argoCdTargetCluster(kubeconfigFlagPath string) {
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

// cniBootstrap installs the Cni Helmchartproxy to the bootstrap cluster
func cniBootstrap(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.CniName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// cniMgmt installs the Cni Helmchartproxy to the mgmt cluster
func cniMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.CniName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// argoAppsMgmt isntalls the ArgoCD Apps Helmchartproxy to the mgmt cluster
func argoAppsMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.ArgoCDAppsMgmt)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// argoEventsMgmt installs the ArgoCD Events Helmchartproxy to the mgmt cluster
func argoEventsMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.ArgoEventsMgmt)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// argoRolloutsMgmt installs the ArgoCD Rollouts Helmchartproxy to the mgmt cluster
func argoRolloutsMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.ArgoRolloutsMgmt)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// argoWorkflowsMgmt installs the ArgoCD Workflows Helmchartproxy to the mgmt cluster
func argoWorkflowsMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.ArgoWorkflowsMgmt)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// argoCDMgmt installs the ArgoCD Helmchartproxy to the mgmt cluster
func argoCDMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.ArgoCDMgmt)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// vaultMgmt installs the Vault Helmchartproxy to the mgmt cluster
func vaultMgmt(kubeconfigFlagPath string) {
	helmchartManifests := filepath.Join(friggDir, vars.VaultName)

	cmd := exec.Command(kubectlPath, "--kubeconfig",
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

// InstallOnBootstrap combined installation on bootstrap cluster
func InstallOnBootstrap() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)
	cniBootstrap(kubeconfigFlagPath)
}

// InstallOnMgmt combined installation on mgmt cluster
func InstallOnMgmt() {
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)
	argoCdTargetCluster(kubeconfigFlagPath)
	cniMgmt(kubeconfigFlagPath)
	argoAppsMgmt(kubeconfigFlagPath)
	argoEventsMgmt(kubeconfigFlagPath)
	argoRolloutsMgmt(kubeconfigFlagPath)
	argoWorkflowsMgmt(kubeconfigFlagPath)
	argoCDMgmt(kubeconfigFlagPath)
	vaultMgmt(kubeconfigFlagPath)
}
