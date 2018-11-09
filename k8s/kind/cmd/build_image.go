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
	"os"

	"github.com/alejandrox1/cool_go/k8s/kind/pkg/build"
	"github.com/spf13/cobra"
)

type buildImageCommandFlags struct {
	Source string
}

func newBuildImageCommand() *cobra.Command {
	flags := &buildImageCommandFlags{}
	cmd := &cobra.Command{
		Use:   "image",
		Short: "build the node image",
		Long:  "build the node image",
		Run: func(cmd *cobra.Command, args []string) {
			runBuildImage(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(&flags.Source, "source", "", "path to node image sources")
	cmd.MarkFlagRequired("source")
	return cmd
}

func runBuildImage(flags *buildImageCommandFlags, cmd *cobra.Command, args []string) {
	ctx := build.NewNodeImageBuildContext()
	ctx.SourceDir = flags.Source
	if err := ctx.Build(); err != nil {
		os.Exit(1)
	}
}
