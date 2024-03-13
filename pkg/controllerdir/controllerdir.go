package controllerdir

import (
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

var ControllerDir string

func FriggControllerDir() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// Setting Tools directory
	ControllerDir = filepath.Join(homedir, vars.FriggDirName, vars.ControllerDir)

	// Check if directory exists and create it if not
	if _, err := os.Stat(ControllerDir); os.IsNotExist(err) {
		println(color.YellowString("Tools directory does not exists, Creating %s\n", ControllerDir))
		if err = os.MkdirAll(ControllerDir, 0755); err != nil {
			println(color.RedString("Error creating directory %s: %v\n", ControllerDir, err))
		}
	} else if err != nil {
		// Handle other potential errors during stat
		println(color.RedString("Error checking directory %s: %v\n", ControllerDir, err))
		os.Exit(1)
	}
}
