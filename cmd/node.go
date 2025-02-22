// Copyright © 2017 Aqua Security Software Ltd. <info@aquasec.com>
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
	"github.com/stuartw1/kube-bench/check"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Run benchmark checks for a Kubernetes node.",
	Long:  `Run benchmark checks for a Kubernetes node.`,
	Run: func(cmd *cobra.Command, args []string) {
		runChecks(check.NODE)
	},
}

func init() {
	nodeCmd.PersistentFlags().StringVarP(&nodeFile,
		"file",
		"f",
		"/node.yaml",
		"Alternative YAML file for node checks",
	)

	RootCmd.AddCommand(nodeCmd)
}
