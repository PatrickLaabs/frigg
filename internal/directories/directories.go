package directories

import (
	"github.com/PatrickLaabs/frigg/internal/vars"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

var (
	friggdir string
)

func workDir() {
	dir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the .frigg working directory: %v\n", err))
		return
	}

	friggdir := filepath.Join(dir, ".frigg")

	// Check if the directory already exists
	if _, err := os.Stat(friggdir); err == nil {
		// Directory already exists, skip creation
		println(color.YellowString("frigg working directory already exists, skipping creation."))
		return
	}

	// Create the directory if it doesn't exist
	if err := os.Mkdir(friggdir, 0750); err != nil {
		println(color.RedString("Error Creating .frigg Directory: %v\n", err))
	}
}

func toolsDir() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// Setting Tools directory
	friggdir = filepath.Join(homedir, vars.FriggDirName, vars.FriggTools)

	// Check if directory exists and create it if not
	if _, err := os.Stat(friggdir); os.IsNotExist(err) {
		println(color.YellowString("Tools directory does not exists, Creating %s\n", friggdir))
		if err = os.MkdirAll(friggdir, 0755); err != nil {
			println(color.RedString("Error creating directory %s: %v\n", friggdir, err))
		}
	} else if err != nil {
		// Handle other potential errors during stat
		println(color.RedString("Error checking directory %s: %v\n", friggdir, err))
		os.Exit(1)
	}
}

func controllerDir() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	// Setting Tools directory
	controllerDir := filepath.Join(homedir, vars.FriggDirName, vars.ControllerDir)

	// Check if directory exists and create it if not
	if _, err := os.Stat(controllerDir); os.IsNotExist(err) {
		println(color.YellowString("Tools directory does not exists, Creating %s\n", controllerDir))
		if err = os.MkdirAll(controllerDir, 0755); err != nil {
			println(color.RedString("Error creating directory %s: %v\n", controllerDir, err))
		}
	} else if err != nil {
		// Handle other potential errors during stat
		println(color.RedString("Error checking directory %s: %v\n", controllerDir, err))
		os.Exit(1)
	}
}

// Create creates friggs working directories .frigg and tools at the users home directory.
func Create() {
	workDir()
	toolsDir()
	controllerDir()
}
