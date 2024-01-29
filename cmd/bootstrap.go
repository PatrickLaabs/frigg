/*
Copyright © 2024 Patrick Laabs <patrick.laabs@me.com>
*/
package cmd

import (
	"fmt"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/kind"
	"github.com/spf13/cobra"
)

// bootstrapCmd represents the create command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Commands to manage KinD LifeCycle",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bootrap called")
		fmt.Println("To learn more about the bootstrap of kind for argohub, run:")
		fmt.Println("  argohub bootstrap --help")
	},
}

func init() {
	bootstrapCmd.AddCommand(
		kind.CreateCmd,
		kind.DeleteCmd,
		kind.KindApiUsageCmd,
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
