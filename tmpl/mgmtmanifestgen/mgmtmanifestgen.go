package mgmtmanifestgen

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
)

func Gen() {
	homedir, _ := os.UserHomeDir()

	friggDirname := ".frigg"

	tmplmanifest := "templates/argohubmgmtclusterManifest.yaml"
	outputmanifestName := "argohubmgmtclusterManifest.yaml"
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

	println(color.GreenString("Successfully written Mgmt Kubernets Manifest"))
}
