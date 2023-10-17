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
	"context"
	"fmt"
	"time"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// saveWeeklyCmd represents the saveWeekly command
var saveWeeklyCmd = &cobra.Command{
	Use:   "save-weekly",
	Short: "Saves the current Spotify Discover Weekly playlist",
	Long: `Saves the current Spotify Discover Weekly playlist.
Note this cannot bring back old Spotify Discover Weekly playlists, it can
only save the current playlist`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		playlistName, _ := cmd.Flags().GetString("name")
		playlistDescription := "Spotify Discover Weekly for " + getPlaylistDate()
		if playlistName == "" {
			playlistName = "Discover Weekly " + getPlaylistDate()
		}
		isPublic, _ := cmd.Flags().GetBool("public")
		isCollaborative, _ := cmd.Flags().GetBool("collaborative")

		ctx := context.Background()
		currentUser, err := conf.Client.CurrentUser(ctx)
		if err != nil {
			helper.LogErrorAndExit(err)
		}
		if deduplicatePlaylist(playlistName, currentUser.User.ID) {
			fmt.Println("Discover Weekly already saved")
			return
		}
		newPlaylist, err := conf.Client.CreatePlaylistForUser(ctx, currentUser.User.ID, playlistName, playlistDescription, isPublic, isCollaborative)
		if err != nil {
			helper.LogErrorAndExit(err)
		}
		searchResult, err := conf.Client.Search(ctx, "Discover Weekly", spotify.SearchTypePlaylist)
		if err != nil {
			helper.LogErrorAndExit(err)
		}
		var discoverPlaylist spotify.ID
		for _, playlist := range searchResult.Playlists.Playlists {
			if playlist.Owner.ID == "spotify" {
				discoverPlaylist = playlist.ID
				break
			}
		}
		discoverPlaylistTracks := func() spotify.PlaylistTrackPage {
			playlistTracks, err := conf.Client.GetPlaylistTracks(ctx, discoverPlaylist)
			if err != nil {
				helper.LogErrorAndExit(err)
			}
			return *playlistTracks
		}
		var discoverPlaylistTrackIDs []spotify.ID
		for _, track := range discoverPlaylistTracks().Tracks {
			discoverPlaylistTrackIDs = append(discoverPlaylistTrackIDs, track.Track.ID)
		}
		conf.Client.AddTracksToPlaylist(ctx, newPlaylist.ID, discoverPlaylistTrackIDs...)
		fmt.Printf("Discover Weekly saved as %s\n", playlistName)
	},
}

func getPlaylistDate() string {
	date := time.Now()
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}
	return fmt.Sprintf("%d/%d/%d", date.Month(), date.Day(), date.Year())
}
func deduplicatePlaylist(playlistName string, user string) bool {
	if conf.Client == nil {
		fmt.Println("Client is not initialized")
		return false
	}
	searchResults, err := conf.Client.Search(context.Background(), playlistName, spotify.SearchTypePlaylist)
	if err != nil {
		helper.LogErrorAndExit(err)
	}
	for _, playlist := range searchResults.Playlists.Playlists {
		if playlist.Owner.ID == user && playlist.Name == playlistName {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(saveWeeklyCmd)
	saveWeeklyCmd.Flags().StringP("name", "n", "", "Custom name for the save playlist")
	saveWeeklyCmd.Flags().BoolP("public", "p", false, "Whether to make the new playlist public")
	saveWeeklyCmd.Flags().BoolP("collaborative", "c", false, "Set the playlist as collaborative")
}
