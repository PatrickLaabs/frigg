package workloadmanifestgen

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func Gen() {
	fmt.Println("Generating Workload Cluster Manifest")

	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"

	tmplmanifest := "templates/workloadcluster.yaml"
	outputmanifestName := "workloadcluster.yaml"
	outputPath := homedir + "/" + friggDirname + "/" + outputmanifestName

	var manifest map[string]interface{}

	yamlFile, err := os.ReadFile(tmplmanifest)
	if err != nil {
		fmt.Printf("Error Reading file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, &manifest)
	if err != nil {
		fmt.Printf("Error unmarshaling yaml: %s", err)
	}

	err = os.WriteFile(outputPath, yamlFile, 0755)
	if err != nil {
		fmt.Printf("Error writing file: %s", err)
	}

	fmt.Println("Successfully written Workload Cluster Kubernets Manifest")
}
