package clusterctlconfig

import (
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type AutoGenerated struct {
	CLUSTERTOPOLOGY string `yaml:"CLUSTER_TOPOLOGY"`
	EXPMACHINEPOOL  string `yaml:"EXP_MACHINE_POOL"`
}

// ClusterctlConfigGen generates clusterctl config to frigg dir
func ClusterctlConfigGen() {
	data := &AutoGenerated{
		CLUSTERTOPOLOGY: "true",
		EXPMACHINEPOOL:  "true",
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marshaling data to yaml: %v\n", err))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newFilePath := filepath.Join(friggDir, vars.ClusterctlConfigName)

	// Write to file
	err = os.WriteFile(newFilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing clusterctl config yaml: %v\n", err))
	}
}
