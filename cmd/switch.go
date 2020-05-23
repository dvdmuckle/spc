/*
Copyright © 2020 David Muckle <dvdmuckle@dvdmuckle>

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
	"github.com/dvdmuckle/goify/helper"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Set device to use for all callbacks",
	Long: `Set the device to use when controlling Spotify playback.
	If this entry is empty, it will default to the currently playing device.
	You can clear the set device entry if the device is no longer active.
	This will also switch playback to the device selected if playback is active,
	and can also switch playback to the already configured device.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func transferPlayback(config *helper.Config) {
	client = &config.Client
}

func init() {
	rootCmd.AddCommand(switchCmd)

	switchCmd.Flags().StringP("set", "d", "", "DeviceID to switch to")
	switchCmd.Flags().BoolP("clear", "c", false, "Clear the current device entry")
	switchCmd.Flags().BoolP("noswitch", "n", false, "Transfer playback to the currently configured device")
	switchCmd.Flags().BoolP("print", "p", false, "Only print the currently configured device")
}
