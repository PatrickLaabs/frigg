package statuscheck

import (
	"context"
	"fmt"
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/fatih/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

// checkDeploymentCondition implements the basis functionality to operate inside the cluster and looks for deployments
func checkDeploymentCondition(clientset *kubernetes.Clientset, namespace string, deployment string) (bool, error) {
	x, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployment, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	for _, condition := range x.Status.Conditions {
		if condition.Type == "Available" && condition.Status == "True" {
			return true, nil
		}
	}
	return false, nil
}

// checkCapiOperator basis implemention for checking the status of the CAPI Operator deployment
func checkCapiOperator(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "capi-operator-system"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionCheckCapiOperator checks for all ClusterAPI Operator deployment conditions on the bootstrap cluster
func ConditionCheckCapiOperator() {
	deployments := []string{"capi-operator-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)
	for {
		go checkCapiOperator(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("CAPI Operator deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// ConditionCheckCapiOperatorMgmt checks for all ClusterAPI Operator deployment conditions on the mgmt cluster
func ConditionCheckCapiOperatorMgmt() {
	deployments := []string{"capi-operator-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)
	for {
		go checkCapiOperator(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("CAPI Operator deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkCertManagers basis implemention for checking the status of the Cert-Manager deployment
func checkCertManagers(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "cert-manager"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionsCertManagers checks the deployment state on the bootstrap cluster
func ConditionsCertManagers() {
	deployments := []string{"cert-manager", "cert-manager-webhook", "cert-manager-cainjector"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	for {
		go checkCertManagers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Cert-Manager deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// ConditionsCertManagersMgmt checks the deployment state on the mgmt cluster
func ConditionsCertManagersMgmt() {
	deployments := []string{"cert-manager", "cert-manager-webhook", "cert-manager-cainjector"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCertManagers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Cert-Manager deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkCapiControllers basis implemention for checking the status of the ClusterAPI Controller deployment
func checkCapiControllers(kubeconfigFlagPath string, deployments map[string][]string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	allAvailable := true
	for namespace, namespaceDeployments := range deployments {
		for _, deployment := range namespaceDeployments {
			available, _ := checkDeploymentCondition(clientset, namespace, deployment)
			if available {
			} else {
				allAvailable = false
			}
		}
	}

	if allAvailable {
		readyChan <- true
	}
}

// ConditionsCapiControllers checks the deployment state on the bootstrap cluster
func ConditionsCapiControllers() {
	//deployments := []string{"capi-controller-manager", "capi-kubeadm-control-plane-controller-manager", "capi-kubeadm-bootstrap-controller-manager"}
	deployments := map[string][]string{
		"capi-system":                       {"capi-controller-manager"},
		"capi-kubeadm-control-plane-system": {"capi-kubeadm-control-plane-controller-manager"},
		"capi-kubeadm-bootstrap-system":     {"capi-kubeadm-bootstrap-controller-manager"},
	}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	for {
		go checkCapiControllers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("ClusterAPI Controller deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// ConditionsCapiControllersMgmt checks the deployment state on the mgmt cluster
func ConditionsCapiControllersMgmt() {
	//deployments := []string{"capi-controller-manager", "capi-kubeadm-control-plane-controller-manager", "capi-kubeadm-bootstrap-controller-manager"}
	deployments := map[string][]string{
		"capi-system":                       {"capi-controller-manager"},
		"capi-kubeadm-control-plane-system": {"capi-kubeadm-control-plane-controller-manager"},
		"capi-kubeadm-bootstrap-system":     {"capi-kubeadm-bootstrap-controller-manager"},
	}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCapiControllers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("ClusterAPI Controller deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkCapdControllers checks the ClusterAPI CAPD Provider conditions
func checkCapdControllers(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "capd-system"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionsCapdControllers checks the deployment state on the bootstrap cluster
func ConditionsCapdControllers() {
	deployments := []string{"capd-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	for {
		go checkCapdControllers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Capd Controller deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// ConditionsCapdControllersMgmt checks the deployment state on the mgmt cluster
func ConditionsCapdControllersMgmt() {
	deployments := []string{"capd-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCapdControllers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Capd Controller deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkCaaphControllers checks the ClusterAPI Helm Addon Controller conditions
func checkCaaphControllers(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "caaph-system"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionsCaaphControllers checks the deployment state on the bootstrap cluster
func ConditionsCaaphControllers() {
	deployments := []string{"caaph-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.BootstrapkubeconfigName)

	for {
		go checkCaaphControllers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Caaph Controller deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// ConditionsCaaphControllersMgmt checks the deployment state on the mgmt cluster
func ConditionsCaaphControllersMgmt() {
	deployments := []string{"caaph-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCaaphControllers(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Caaph Controller deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkCni checks CNI deployment conditions
func checkCni(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "calico-system"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionsCni checks the deployment state on the mgmt cluster
func ConditionsCni() {
	deployments := []string{"calico-kube-controllers", "calico-typha"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCni(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("CNI deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkCoreDns checks Core DNS deployment conditions
func checkCoreDns(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "kube-system"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionCoreDnsMgmt checks the deployment state on the mgmt cluster
func ConditionCoreDnsMgmt() {
	deployments := []string{"coredns"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCoreDns(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Core DNS deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// checkTigerOperator checks Tiger-Operator deployment conditions
func checkTigerOperator(kubeconfigFlagPath string, deployments []string, readyChan chan bool) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigFlagPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "kube-system"
	allAvailable := true
	for _, deployment := range deployments {
		available, _ := checkDeploymentCondition(clientset, namespace, deployment)
		if available {
		} else {
			allAvailable = false
		}
	}
	if allAvailable {
		readyChan <- true
	}
}

// ConditionTigerOperatorMgmt checks the deployment state on the mgmt cluster
func ConditionTigerOperatorMgmt() {
	deployments := []string{"tigera-operator"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkTigerOperator(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("Tigera Operator deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}

// ConditionCniWorkload checks the CNI Installation on the mgmt cluster before continuing
func ConditionCniWorkload() {
	deployments := []string{"capi-kubeadm-control-plane-controller-manager"}
	readyChan := make(chan bool) // Create a channel for readiness signal
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)
	kubeconfigFlagPath := filepath.Join(friggDir, vars.ManagementKubeconfigName)

	for {
		go checkCni(kubeconfigFlagPath, deployments, readyChan)

		// Wait for readiness signal or timeout
		select {
		case <-readyChan:
			fmt.Println("CNI deployments are available!")
			return // Exit the loop
		case <-time.After(5 * time.Second):
		}
	}
}
