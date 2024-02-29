package mgmtmanifestgen

import (
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
		println(color.RedString("Error Reading file: %v\n", err))
	}

	err = yaml.Unmarshal(yamlFile, &manifest)
	if err != nil {
		println(color.RedString("Error unmarshaling yaml: %v\n", err))
	}

	err = os.WriteFile(outputPath, yamlFile, 0755)
	if err != nil {
		println(color.RedString("Error writing file: %v\n", err))
	}

	println(color.GreenString("Successfully written Mgmt Kubernets Manifest"))
}
