/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// landscapeCmd represents the landscape command
var landscapeCmd = &cobra.Command{
	Use:   "landscape",
	Short: "Generates a system landscape diagram",
	Long:  `Generates a diagram of root actors and elements, and the associations between them.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Generate the diagram
		diagram, err = arch.LandscapeView()
	},
}

func init() {
	generateCmd.AddCommand(landscapeCmd)
}