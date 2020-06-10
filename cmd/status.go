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

	"github.com/golang/glog"
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

	%[1]a - artist
	%[1]t - track
	%[1]b - album
	%[1]p - playing
	
	If a song has multiple artists, you can specify the upper limit of artists
	to display with %X[1]a, where X is the number of artists to print, separated
	by commas.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, err := conf.Client.PlayerCurrentlyPlaying()
		if err != nil {
			glog.Fatal(err)
		}
		statusFmt, _ := cmd.Flags().GetString("format")
		if err != nil {
			glog.Fatal(err)
		}
		if statusFmt != "" {
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
	case 'b':
		fmt.Fprint(state, stat.Item.Album.Name)
	case 't':
		fmt.Fprint(state, stat.Item.SimpleTrack.Name)
	case 'p':
		fmt.Fprint(state, stat.Playing)
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
