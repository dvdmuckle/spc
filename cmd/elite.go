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
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
)

type MusicEvent struct {
	Timestamp  string `json:"timestamp"`
	Event      string `json:"event`
	MusicTrack string `json:"MusicTrack`
}

// eliteCmd represents the elite command
var eliteCmd = &cobra.Command{
	Use:   "elite",
	Short: "A companion subcommand for Elite Dangerous",
	Long: `A companion subcommand for the game Elite Dangerous. Will auto select
	what playlist to play based on in-game events and user configuration. If
	all of this sounds alien to you, safely ignore this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)

		// Create new watcher.
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			helper.LogErrorAndExit(err)
		}
		defer watcher.Close()

		var lastSeenTrack string
		var currentlyPlayingTrack string

		// Start listening for events.
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Has(fsnotify.Write) {
						if regexp.MustCompile(`Journal\.\d{4}\-\d{2}\-\d{2}T\d{6}\.\d{2}\.log$`).MatchString(event.Name) {
							//TODO: I hate that we read the whole file every time we want to get just the latest pertinent line
							//We should not do this
							journalFile, err := os.Open(event.Name)
							if err != nil {
								helper.LogErrorAndExit(err)
							}
							defer journalFile.Close()
							scanner := bufio.NewScanner(journalFile)
							var journalEvent MusicEvent
							for scanner.Scan() {
								var tempEvent MusicEvent
								json.Unmarshal([]byte(scanner.Text()), &tempEvent)
								if tempEvent.Event == "Music" {
									journalEvent = tempEvent
								}
							}
							if !(lastSeenTrack == journalEvent.MusicTrack) {
								fmt.Println("New track detected: " + journalEvent.MusicTrack)
								lastSeenTrack = journalEvent.MusicTrack
								playlistToSearch := viper.GetString("elite." + journalEvent.MusicTrack)
								if playlistToSearch != "" && lastSeenTrack != currentlyPlayingTrack {
									currentlyPlayingTrack = lastSeenTrack
									ctx := context.Background()
									//searchResult, err := conf.Client.Search(ctx, playlistToSearch, spotify.SearchTypePlaylist)
									//if err != nil {
									//	helper.LogErrorAndExit(err)
									//}
									//opts.PlaybackContext = &searchResult.Playlists.Playlists[0].URI
									var opts spotify.PlayOptions
									opts.DeviceID = &conf.DeviceID
									playlistString := spotify.URI(playlistToSearch)
									opts.PlaybackContext = &playlistString
									fmt.Println("Found something to play for " + currentlyPlayingTrack)
									if err := conf.Client.PlayOpt(ctx, &opts); err != nil {
										helper.LogErrorAndExit(err)
									}
								}
							}
						}
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Println("error:", err)
				}
			}
		}()

		// Add a path.
		err = watcher.Add("/home/dvdmuckle/.steam/steam/steamapps/compatdata/359320/pfx/drive_c/users/steamuser/Saved Games/Frontier Developments/Elite Dangerous")
		if err != nil {
			log.Fatal(err)
		}

		// Block main goroutine forever.
		<-make(chan struct{})
	},
}

func init() {
	rootCmd.AddCommand(eliteCmd)
}
