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

	"github.com/alejandrox1/cool_go/k8s/kind/pkg/cluster"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a cluster",
		Long:  "Create a kubernetes cluster",
		Run:   runCreate,
	}
}

func runCreate(cmd *cobra.Command, args []string) {
	config := cluster.NewConfig("")
	ctx := cluster.NewContext(config)
	if err := ctx.Create(); err != nil {
		glog.Errorf("Failed to create cluster: %v", err)
		os.Exit(1)
	}
}
