/*
Copyright © 2020 David Muckle <dvdmuckle@dvdmuckle.xyz>

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

	"github.com/dvdmuckle/goify/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start Spotify playback",
	Long: `Will start Spotify playback on the device most
	recently playing music.`,
	Run: func(cmd *cobra.Command, args []string) {
		if conf.Client == (spotify.Client{}) {
			fmt.Println("Please run goify auth first to login")
			os.Exit(1)
		}
		helper.Play(&conf)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
