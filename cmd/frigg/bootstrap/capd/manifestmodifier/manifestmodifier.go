package manifestmodifier

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ModifyMgmt() {
	fmt.Println("Modifying manifest, and adding Helmchartproxies labels..")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	mgmtClusterManifestPath := homedir + "/" + argohubDirName + "/" + "argohubmgmtclusterManifest.yaml"

	yamlFile, err := os.ReadFile(mgmtClusterManifestPath)
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		os.Exit(1)
	}

	var data map[interface{}]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		fmt.Println("Error unmarshaling YAML data:", err)
		os.Exit(1)
	}

	metadata := data["metadata"].(map[string]interface{})
	_, hasLabels := metadata["labels"]
	if !hasLabels {
		// "labels" field doesn't exist, create it
		metadata["labels"] = make(map[string]interface{})
	}
	labels := metadata["labels"].(map[interface{}]interface{})
	labels["argocd-hub-Enabled"] = "default"
	labels["argocd-hub-appsEnabled"] = "default"
	labels["argo-events-Enabled"] = "default"
	labels["argo-rollouts-Enabled"] = "default"
	labels["argo-workflows-Enabled"] = "default"
	labels["vaultmgmtEnabled"] = "default"

	modifiedYAML, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println("Error marshaling YAML data:", err)
		os.Exit(1)
	}

	err = os.WriteFile(mgmtClusterManifestPath, modifiedYAML, 0750)
	if err != nil {
		fmt.Println("Error writing modified YAML file:", err)
		os.Exit(1)
	}

	fmt.Println("YAML file modified successfully!")

}

func Modifyworkload() {
	fmt.Println("Modifying manifest, and adding Helmchartproxies labels..")

	homedir, _ := os.UserHomeDir()

	argohubDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	// /home/patricklaabs/.frigg/frigg-cluster.kubeconfig
	kubeconfigFlagPath := homedir + "/" + argohubDirName + "/" + kubeconfigName

	workloadClusterManifestPath := homedir + "/" + argohubDirName + "/" + "gened-Manifest.yaml"

	yamlFile, err := os.ReadFile(workloadClusterManifestPath)
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		os.Exit(1)
	}

	var data map[interface{}]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		fmt.Println("Error unmarshaling YAML data:", err)
		os.Exit(1)
	}

	metadata := data["metadata"].(map[interface{}]interface{})
	labels := metadata["labels"].(map[interface{}]interface{})
	labels["argocd-hub-workloadclusters-apps-Enabled"] = "default"
	labels["argocd-workloadclusters-clusters-Enabled"] = "default"

	modifiedYAML, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println("Error marshaling YAML data:", err)
		os.Exit(1)
	}

	err = os.WriteFile(workloadClusterManifestPath, modifiedYAML, 0750)
	if err != nil {
		fmt.Println("Error writing modified YAML file:", err)
		os.Exit(1)
	}

	fmt.Println("YAML file modified successfully!")

}
