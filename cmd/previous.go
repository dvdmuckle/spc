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
	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// previousCmd represents the previous command
var previousCmd = &cobra.Command{
	Use:   "previous",
	Short: "Goes back to the previously playing track",
	Long:  `Goes back to the previously playing track. Will use the currently configured device.`,
	Run: func(cmd *cobra.Command, args []string) {
		var opts spotify.PlayOptions
		helper.SetClient(&conf, cfgFile)
		opts.DeviceID = &conf.DeviceID
		//We want to go to the last song playing, but
		//spotify.Previous() will rewind the current song
		//unless the current song is close to the beginning
		//of playback, so we seek to zero here before calling
		//spotify.Previous()
		conf.Client.SeekOpt(0, &opts)
		conf.Client.PreviousOpt(&opts)
	},
	Aliases: []string{"prev"},
}

func init() {
	rootCmd.AddCommand(previousCmd)
}
