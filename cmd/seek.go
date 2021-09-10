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
	"strconv"
	"strings"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

var seekCmd = &cobra.Command{
	Use:   "seek",
	Args:  cobra.ExactArgs(1),
	Short: "Seek to a specific position in the currently playing song from Spotify",
	Long: `Seek to a specific position in the currently playing song from Spotify. This command requires
exactly one argument, either a number between 0 and the length of the song in seconds, or a timestamp in
the form of minutes:seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		var position int
		if strings.Contains(args[0], ":") {
			var (
				minutes int
				seconds int
			)
			_, err := fmt.Sscanf(args[0], "%d:%d", &minutes, &seconds)
			if err != nil {
				fmt.Println("Timestamp must be numbers in the form of minutes:seconds")
				os.Exit(1)
			}
			position = minutes*60 + seconds
		} else {
			var err error
			position, err = strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Passed value for seconds must be an integer.")
				os.Exit(1)
			}
		}

		currentlyPlaying, err := conf.Client.PlayerCurrentlyPlaying()
		if err != nil {
			glog.Fatal(err)
		}

		if currentlyPlaying.Item == nil {
			fmt.Println("Could not obtain the currently playing song.")
			os.Exit(1)
		}

		duration := currentlyPlaying.Item.Duration / 1000
		if position > duration {
			fmt.Printf(
				"The seek position must be at or under the duration of the currently playing song (%d seconds, or %d:%d).\n",
				duration, duration/60, duration%60)
			os.Exit(1)
		}

		err = conf.Client.SeekOpt(position*1000, &spotify.PlayOptions{DeviceID: &conf.DeviceID})
		if err != nil {
			glog.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(seekCmd)
}
