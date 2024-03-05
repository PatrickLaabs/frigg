package download

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

func init() {
	FriggWorkingDir()
}

var friggDir string

func FriggWorkingDir() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// Setting Tools directory
	friggDir = filepath.Join(homedir, vars.FriggDirName, vars.FriggTools)

	// Check if directory exists and create it if not
	if _, err := os.Stat(friggDir); os.IsNotExist(err) {
		println(color.YellowString("Tools directory does not exists, Creating %s\n", friggDir))
		err = os.MkdirAll(friggDir, 0755) // Create directory with permissions 0755
		if err != nil {
			println(color.RedString("Error creating directory %s: %v\n", friggDir, err))
		}
	} else if err != nil {
		// Handle other potential errors during stat
		println(color.RedString("Error checking directory %s: %v\n", friggDir, err))
		os.Exit(1)
	}
}

//- [kind](https://formulae.brew.sh/formula/kind#default)
//- [k9s](https://formulae.brew.sh/formula/k9s#default)
//- [Docker](https://www.docker.com/products/docker-desktop/)
//- [clusterctl](https://formulae.brew.sh/formula/clusterctl#default)
//- [kubectl](https://formulae.brew.sh/formula/kubernetes-cli#default)
//- [github cli](https://formulae.brew.sh/formula/gh#default)
//- [jq](https://formulae.brew.sh/formula/jq#default)
//- [helm]()

func Helm() {
	data := "Bla"

	filePath := filepath.Join(friggDir, "test.yaml")
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		println(color.RedString("error on writing helm tarball to working directory: %v\n", err))
	}
}
