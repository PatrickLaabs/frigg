/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY frigg, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package frigg implements the root frigg cobra command, and the cli Main()
package frigg

import (
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap"
	"github.com/PatrickLaabs/frigg/cmd/frigg/prepare"
	"io"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/cmd/frigg/delete"
	"github.com/PatrickLaabs/frigg/cmd/frigg/version"
	"github.com/PatrickLaabs/frigg/pkg/log"
)

type flagpole struct {
	LogLevel  string
	Verbosity int32
	Quiet     bool
}

// NewCommand returns a new cobra.Command implementing the root command for frigg
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	flags := &flagpole{}
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "frigg",
		Short: "frigg provisions kubernetes cluster with capi and gitops in no-time",
		Long:  "frigg provisions kubernetes cluster with capi and gitops in no-time",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runE(logger, flags, cmd)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.Version(),
	}
	c.SetOut(streams.Out)
	c.SetErr(streams.ErrOut)
	c.PersistentFlags().StringVar(
		&flags.LogLevel,
		"loglevel",
		"",
		"DEPRECATED: see -v instead",
	)
	c.PersistentFlags().Int32VarP(
		&flags.Verbosity,
		"verbosity",
		"v",
		0,
		"info log verbosity, higher value produces more output",
	)
	c.PersistentFlags().BoolVarP(
		&flags.Quiet,
		"quiet",
		"q",
		false,
		"silence all stderr output",
	)
	// add all top level subcommands
	c.AddCommand(delete.NewCommand(logger))
	c.AddCommand(version.NewCommand(logger, streams))
	c.AddCommand(bootstrap.NewCommand(logger, streams))
	c.AddCommand(prepare.NewCommand())
	return c
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
