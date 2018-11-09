// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"flag"

	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kind",
		Short: "kind is a tool for managing local multi-node kubernetes clusters",
		Long: `kind creates and manages local multi-node kubernetes clusters
using Docker containers`,
	}

	// Add all top-level commands.
	cmd.AddCommand(newBuildCommand())
	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newDeleteCommand())
	return cmd
}

// Run runs the 'kind' root command.
func Run() error {
	// Trick to avoid glog's 'logging before flag.Parse' warning
	flag.CommandLine.Parse([]string{})
	// glog logs to files by default.
	flag.Set("logtostderr", "true")

	rootCmd := newRootCommand()
	// glog registers global flags on flag.CommandLine
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)

	// Execute command now.
	return rootCmd.Execute()
}
