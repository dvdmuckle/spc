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
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell autocompletion",
	Long: `Generates shell autocompletion

The following shells can have autocompletion generated:

bash
zsh
fish
powershell

For fish, the flag --fish-description can be toggled to includes descriptions in the autocomplete`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Please provide a shell for which to generate autocompletion")
			os.Exit(1)
		} else if len(args) < 1 {
			fmt.Println("Please provide only one shell for which to generate autocompletion")
		}
		if args[0] == "bash" {
			rootCmd.GenBashCompletion(os.Stdout)
			return
		}
		if args[0] == "zsh" {
			rootCmd.GenZshCompletion(os.Stdout)
			return
		}
		if args[0] == "powershell" {
			rootCmd.GenPowerShellCompletion(os.Stdout)
			return
		}
		if args[0] == "fish" {
			description, _ := cmd.Flags().GetBool("fish-description")
			rootCmd.GenFishCompletion(os.Stdout, description)
			return
		}
		fmt.Println(args)
	},
	ValidArgs: []string{"bash", "zsh", "powershell", "fish"},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().Bool("fish-description", false, "Whether to include description for fish autocompletion")
}
