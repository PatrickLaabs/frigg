package toolsdir

import (
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

var FriggDir string

func FriggWorkingDir() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// Setting Tools directory
	FriggDir = filepath.Join(homedir, vars.FriggDirName, vars.FriggTools)

	// Check if directory exists and create it if not
	if _, err := os.Stat(FriggDir); os.IsNotExist(err) {
		println(color.YellowString("Tools directory does not exists, Creating %s\n", FriggDir))
		if err = os.MkdirAll(FriggDir, 0755); err != nil {
			println(color.RedString("Error creating directory %s: %v\n", FriggDir, err))
		}
	} else if err != nil {
		// Handle other potential errors during stat
		println(color.RedString("Error checking directory %s: %v\n", FriggDir, err))
		os.Exit(1)
	}
}
