package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "argohub",
	Short: "argohub is a very fast gitops bootstrap toolset",
	Long: `A Fast and Flexible bootstrap generator for gitops built with
                love by Patrick Laabs in go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To learn more about argohub, run:")
		fmt.Println("  argohub help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.SilenceUsage = true
	rootCmd.AddCommand(
		bootstrapCmd,
		versionCmd,
	)
}
