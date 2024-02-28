package postbootstrap

import (
	"github.com/fatih/color"
	"os/exec"
)

func DeleteBootstrapcluster() {
	println(color.YellowString("deleting bootstrap cluster.."))

	cmd := exec.Command("kind", "delete", "clusters",
		"bootstrapcluster",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error on deletion of bootstrap cluster: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}
}
