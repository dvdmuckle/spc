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
	"regexp"
	"strconv"
	"time"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

type Status spotify.CurrentlyPlaying

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the currently playing song from Spotify",
	Long: `Get the currently playing song from Spotify

A format string can be passed with --format or -f to specify what
the status printout should look like. The following fields are available:

%a - Artist
%t - Track
%b - Album
%f - Playing
%e ️- ▶ or ⏸️
%s - Play progress

If a song has multiple artists, you can specify the upper limit of artists
to display with %Xa, where X is the number of artists to print, separated
by commas.

If there is no currently playing song on Spotify, regardless of format argument
the command will return an empty string. This may happen if Spotify is paused
for an extended period of time`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		status, err := conf.Client.PlayerCurrentlyPlaying()
		if err != nil {
			helper.LogErrorAndExit(err)
		}

		statusFmt, _ := cmd.Flags().GetString("format")

		if status.Item == nil {
			fmt.Print("")
			return
		}
		if statusFmt != "" {
			reg := regexp.MustCompile(`%([\d])?([atbf])`)
			statusFmt = reg.ReplaceAllString(statusFmt, `%$1[1]$2`)
			toFmt := Status(*status)
			fmt.Printf(statusFmt+"\n", toFmt)
		} else {
			fmt.Printf("%s - %s\n", status.Item.SimpleTrack.Artists[0].Name, status.Item.SimpleTrack.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("format", "f", "", "Format string for formatting the status")
}

//Format implements Formatter for Spotify status
func (stat Status) Format(state fmt.State, verb rune) {
	switch verb {
	case 'f':
		fmt.Fprint(state, strconv.FormatBool(stat.Playing))
	case 'e':
		if stat.Playing {
			fmt.Fprint(state, "▶️")
		} else {
			fmt.Fprint(state, "⏸️")
		}
	case 'b':
		fmt.Fprint(state, stat.Item.Album.Name)
	case 't':
		fmt.Fprint(state, stat.Item.SimpleTrack.Name)
	case 's':
		trackProgress := time.Duration(int64(stat.Progress) * int64(time.Millisecond))
		trackLength := stat.Item.TimeDuration()
		trackProgressMinutes := fmt.Sprintf("%.0f", trackProgress.Truncate(time.Minute).Minutes())
		trackProgressSeconds := fmt.Sprintf("%02d", int64(trackProgress.Seconds())%60)
		trackLengthMinutes := fmt.Sprintf("%.0f", trackLength.Truncate(time.Minute).Minutes())
		trackLengthSeconds := fmt.Sprintf("%02d", int64(trackLength.Seconds())%60)
		fmt.Fprint(state, trackProgressMinutes+":"+trackProgressSeconds+"/"+trackLengthMinutes+":"+trackLengthSeconds)
	case 'a':
		wid, set := state.Width()
		if set {
			for i, artist := range stat.Item.SimpleTrack.Artists {
				//If we are printing the last artist, we don't want a comma
				if i == wid-1 || i == len(stat.Item.SimpleTrack.Artists)-1 {
					fmt.Fprint(state, artist.Name)
					break
				} else {
					fmt.Fprint(state, artist.Name+", ")
				}
			}
		} else {
			fmt.Fprint(state, stat.Item.SimpleTrack.Artists[0].Name)
		}
	}
}
