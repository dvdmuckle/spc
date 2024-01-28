/*
Copyright © 2024 David Muckle <dvdmuckle@dvdmuckle.xyz>

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
	"os"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generates docs fpr spc",
	Long: `Generates docs for spc.
This command is mostly used for automation purposes, but can be used to generate
either man page or markdown documentation. The first argument is which
kind of documentation to generate, either man or markdown. The second is the path for the
generated docs. If the path does not exist, it will be created.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || args[0] != "man" && args[0] != "markdown" {
			fmt.Println("Please supply a doc type, either man or markdown")
			os.Exit(1)
		}
		if len(args) != 2 {
			fmt.Println("Only two args are valid, the doc type and the path")
			os.Exit(1)
		}
		if err := os.MkdirAll(args[1], 0755); err != nil {
			helper.LogErrorAndExit("Error creating docs path: ", err)
		}
		if genTag, err := cmd.Flags().GetBool("gen-tags"); genTag && err == nil {
			rootCmd.DisableAutoGenTag = false
		} else if err != nil {
			helper.LogErrorAndExit(err)
		}
		if args[0] == "man" {
			err := doc.GenManTree(rootCmd, nil, args[1])
			if err != nil {
				helper.LogErrorAndExit(err)
			}

		} else if args[0] == "markdown" {
			err := doc.GenMarkdownTree(rootCmd, args[1])
			if err != nil {
				helper.LogErrorAndExit(err)
			}
		}
	},
	ValidArgs: []string{"man", "markdown"},
}

func init() {
	rootCmd.AddCommand(docsCmd)
	docsCmd.Flags().Bool("gen-tags", false, "Add autogentags to generated docs")
}
