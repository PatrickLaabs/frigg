/*
Copyright © 2024 Patrick Laabs <patrick.laabs@me.com>
*/
package kind

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// CreateCmd represents the localBootstrap command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creation of KinD Cluster",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating in Docker... called")
		clusterName := "argohub-cluster"

		// Command to create the kind cluster
		kindCmd := exec.Command("kind", "create", "cluster",
			"--name", clusterName,
		)

		// Set the output to os.Stdout and os.Stderr
		kindCmd.Stdout = os.Stdout
		kindCmd.Stderr = os.Stderr

		// Run the command
		err := kindCmd.Run()
		if err != nil {
			fmt.Printf("Error creating kind cluster: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully created kind cluster '%s'\n", clusterName)
	},
}
