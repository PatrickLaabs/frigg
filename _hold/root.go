package _hold

import (
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/build"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/create"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/export"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/get"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/load"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cmd/kind/version"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"
	"github.com/spf13/cobra"
	"io"
)

//var rootCmd = &cobra.Command{
//	Use:   "argohub",
//	Short: "argohub is a very fast gitops bootstrap toolset",
//	Long: `A Fast and Flexible bootstrap generator for gitops built with
//                love by Patrick Laabs in go.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("To learn more about argohub, run:")
//		fmt.Println("  argohub help")
//	},
//}
//
//func Execute() {
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//}
//
//func init() {
//	cobra.OnInitialize()
//	rootCmd.SilenceUsage = true
//	rootCmd.AddCommand(
//		bootstrapCmd,
//		versionCmd,
//	)
//}

type flagpole struct {
	LogLevel  string
	Verbosity int32
	Quiet     bool
}

// NewCommand returns a new cobra.Command implementing the root command for kind
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "kind",
		Short: "kind is a tool for managing local Kubernetes clusters",
		Long:  "kind creates and manages local Kubernetes clusters using Docker container 'nodes'",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runE(logger, flags, cmd)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.Version(),
	}
	cmd.SetOut(streams.Out)
	cmd.SetErr(streams.ErrOut)
	cmd.PersistentFlags().StringVar(
		&flags.LogLevel,
		"loglevel",
		"",
		"DEPRECATED: see -v instead",
	)
	cmd.PersistentFlags().Int32VarP(
		&flags.Verbosity,
		"verbosity",
		"v",
		0,
		"info log verbosity, higher value produces more output",
	)
	cmd.PersistentFlags().BoolVarP(
		&flags.Quiet,
		"quiet",
		"q",
		false,
		"silence all stderr output",
	)
	// add all top level subcommands
	cmd.AddCommand(build.NewCommand(logger, streams))
	cmd.AddCommand(create.NewCommand(logger, streams))
	//cmd.AddCommand(delete.NewCommand(logger, streams))
	cmd.AddCommand(export.NewCommand(logger, streams))
	cmd.AddCommand(get.NewCommand(logger, streams))
	cmd.AddCommand(version.NewCommand(logger, streams))
	cmd.AddCommand(load.NewCommand(logger, streams))
	return cmd
}

func runE(logger log.Logger, flags *flagpole, command *cobra.Command) error {
	// handle limited migration for --loglevel
	setLogLevel := command.Flag("loglevel").Changed
	setVerbosity := command.Flag("verbosity").Changed
	if setLogLevel && !setVerbosity {
		switch flags.LogLevel {
		case "debug":
			flags.Verbosity = 3
		case "trace":
			flags.Verbosity = 2147483647
		}
	}
	// normal logger setup
	if flags.Quiet {
		// NOTE: if we are coming from app.Run handling this flag is
		// redundant, however it doesn't hurt, and this may be called directly.
		maybeSetWriter(logger, io.Discard)
	}
	maybeSetVerbosity(logger, log.Level(flags.Verbosity))
	// warn about deprecated flag if used
	if setLogLevel {
		if cmd.ColorEnabled(logger) {
			logger.Warn("\x1b[93mWARNING\x1b[0m: --loglevel is deprecated, please switch to -v and -q!")
		} else {
			logger.Warn("WARNING: --loglevel is deprecated, please switch to -v and -q!")
		}
	}
	return nil
}

// maybeSetWriter will call logger.SetWriter(w) if logger has a SetWriter method
func maybeSetWriter(logger log.Logger, w io.Writer) {
	type writerSetter interface {
		SetWriter(io.Writer)
	}
	v, ok := logger.(writerSetter)
	if ok {
		v.SetWriter(w)
	}
}

// maybeSetVerbosity will call logger.SetVerbosity(verbosity) if logger
// has a SetVerbosity method
func maybeSetVerbosity(logger log.Logger, verbosity log.Level) {
	type verboser interface {
		SetVerbosity(log.Level)
	}
	v, ok := logger.(verboser)
	if ok {
		v.SetVerbosity(verbosity)
	}
}
