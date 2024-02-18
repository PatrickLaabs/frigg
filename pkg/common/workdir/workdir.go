package workdir

import (
	"fmt"
	"os"
)

func CreateDir() {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	err = os.Mkdir(dir+"/.frigg", 0750)
	if err != nil {
		fmt.Printf("Error Creating .frigg Directory %v\n", err)
	}
}
