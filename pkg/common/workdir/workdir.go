package workdir

import (
	"github.com/fatih/color"
	"os"
)

func CreateDir() {
	dir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the .frigg working directory: %v\n", err))
	}

	err = os.Mkdir(dir+"/.frigg", 0750)
	if err != nil {
		println(color.RedString("Error Creating .frigg Directory %v\n", err))
	}
}
