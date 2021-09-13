/*
Copyright © 2021 David Muckle <dvdmuckle@dvdmuckle.xyz>

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
	"fmt"

	"runtime/debug"

	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of spc",
	Long:  `Print the version of spc`,
	Run: func(cmd *cobra.Command, args []string) {
		if version == "" {
			mainModule, ok := debug.ReadBuildInfo()
			if !ok {
				fmt.Println("Error getting version")
			}
			version = mainModule.Main.Version
		}
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
