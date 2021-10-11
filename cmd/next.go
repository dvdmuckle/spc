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
	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Skips the track currently playing",
	Long:  `Skips the track currently playing. Will use the currently configured device.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf, verboseErrLog)
		conf.Client.NextOpt(&spotify.PlayOptions{DeviceID: &conf.DeviceID})
	},
	Aliases: []string{"skip"},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
