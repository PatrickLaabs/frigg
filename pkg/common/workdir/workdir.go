package workdir

import (
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

func CreateDir() {
	println(color.YellowString("Creating Working directory"))
	dir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the .frigg working directory: %v\n", err))
		return
	}

	friggDir := filepath.Join(dir, ".frigg")

	// Check if the directory already exists
	if _, err := os.Stat(friggDir); err == nil {
		// Directory already exists, skip creation
		println(color.YellowString("frigg working directory already exists, skipping creation."))
		return
	}

	// Create the directory if it doesn't exist
	if err := os.Mkdir(friggDir, 0750); err != nil {
		println(color.RedString("Error Creating .frigg Directory: %v\n", err))
	}
}
