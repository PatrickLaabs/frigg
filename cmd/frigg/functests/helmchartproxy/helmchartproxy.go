package helmchartproxy

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/common/sshkey"
	"github.com/PatrickLaabs/frigg/tmpl/helmchartsproxies"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "hc",
		Short: "modify hc func",
		Long:  "modify hc func",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("modify hc func")
			homedir, _ := os.UserHomeDir()
			friggDirName := ".frigg"
			friggDir := homedir + "/" + friggDirName

			username, err := retrieveGithubUserEnv()
			if err != nil {
				fmt.Println("Error retrieving github username:", err)
				os.Exit(1)
			}

			sshkey.KeypairGen()

			filePath := "templates/helmchartproxies/mgmt-argocd_ssh.yaml"
			newFile := "mgmt-argocd.yaml"
			newfilePath := friggDir + "/" + newFile

			sshprivatekeyPath := friggDir + "/" + "frigg-sshkeypair_gen"
			sshprivatekey, err := os.ReadFile(sshprivatekeyPath)
			sshprivatekeyTrimmed := strings.TrimSuffix(string(sshprivatekey), "\n")

			fmt.Println(sshprivatekeyTrimmed)
			err = helmchartsproxies.MGmtArgoCDReplacementTest(filePath, newfilePath, username, sshprivatekeyTrimmed)
			if err != nil {
				println(color.RedString("error on string replacement for sshkeypair: %v\n", err))
			}
			return err
		},
	}
	return c
}

// retrieveGithubUserEnv retrieves and reads the os.Env variables needed for further preperation
// GITHUB_USER
func retrieveGithubUserEnv() (string, error) {
	// Get GITHUB_USERNAME environment var
	var username string

	if os.Getenv("GITHUB_USERNAME") == "" {
		println(color.RedString("Missing Github Username, please set it. Exiting now."))
		os.Exit(1)
	} else {
		username = os.Getenv("GITHUB_USERNAME")
	}

	return username, nil
}
