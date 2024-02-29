package sshkeygen

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/pkg/common/sshkey"
	"github.com/PatrickLaabs/frigg/pkg/log"
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "sshkeygen",
		Short: "sshkeygen",
		Long:  "sshkeygen",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("sshkeygen")

			sshkey.KeypairGen()

			return nil
		},
	}
	return c
}
