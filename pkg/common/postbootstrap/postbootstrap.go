package postbootstrap

import (
	"fmt"
	"os/exec"
)

func DeleteBootstrapcluster() {
	fmt.Println("deleting bootstrap cluster..")

	cmd := exec.Command("kind", "delete", "clusters",
		"bootstrapcluster",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))

}
