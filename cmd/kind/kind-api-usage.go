/*
Copyright © 2024 Patrick Laabs <patrick.laabs@me.com>
*/
package kind

import (
	"fmt"
	"github.com/spf13/cobra"
)

// KindApiUsageCmd represents the localBootstrap command
var KindApiUsageCmd = &cobra.Command{
	Use:   "kindApiUsageCmd",
	Short: "Using Go Pkg of KinD",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("KindApiUsageCmd called")
	},
}
