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
	"context"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
)

// shuffleCmd represents the shuffle command
var shuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Toggle Spotify shuffle",
	Long:  `Will toggle shuffle on or off depending on its previous state`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		ctx := context.Background()
		forceSetShuffle, _ := cmd.Flags().GetBool("set")
		if forceSetShuffle {
			conf.Client.Shuffle(ctx, true)
			return
		}
		state, err := conf.Client.PlayerState(ctx)
		if err != nil {
			helper.LogErrorAndExit("Error getting player state: ", err)
		}
		conf.Client.Shuffle(ctx, !state.ShuffleState)
	},
}

func init() {
	rootCmd.AddCommand(shuffleCmd)
	shuffleCmd.Flags().BoolP("set", "s", false, "Set the shuffle state to on regardless of current state")
}
