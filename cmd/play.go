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
	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start Spotify playback",
	Long:  `Will start Spotify playback on the device most recently playing music.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		helper.Play(&conf)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
