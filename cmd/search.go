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
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for and play a track, album, playlist, or artist",
	Long: `Search takes two arguments: the search type, and the query.
Search type can be an album, a track, a playlist, or an artist, with the rest of the arguments
making up the search query. For example:

	spc search album moving pictures
	spc search track tom sawyer
	spc search playlist prog monsters
	spc search artist rush

If a track is queried for, additional similar songs will be queued up.

More advanced options are available for the search query. For this,
please see https://pkg.go.dev/github.com/zmb3/spotify?tab=doc#Client.Search`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		var searchType string
		if len(args) < 2 {
			fmt.Println("Please search for a track, album, playlist, or artist")
			os.Exit(1)
		}
		if args[0] == "track" || args[0] == "album" || args[0] == "playlist" || args[0] == "artist" {
			searchType = args[0]
		} else {
			fmt.Println("Please search for a track, album, playlist, or artist")
			os.Exit(1)
		}
		searchTerm := strings.Join(args[1:], " ")
		var searchResults *spotify.SearchResult
		var err error
		ctx := context.Background()

		switch searchType {
		case "track":
			searchResults, err = conf.Client.Search(ctx, searchTerm, spotify.SearchTypeTrack)
		case "album":
			searchResults, err = conf.Client.Search(ctx, searchTerm, spotify.SearchTypeAlbum)
		case "playlist":
			searchResults, err = conf.Client.Search(ctx, searchTerm, spotify.SearchTypePlaylist)
		case "artist":
			searchResults, err = conf.Client.Search(ctx, searchTerm, spotify.SearchTypeArtist)
		}
		if err != nil {
			helper.LogErrorAndExit(err)
		}
		toPlay := fuzzySearchResults(*searchResults, searchType)
		var opts spotify.PlayOptions
		opts.DeviceID = &conf.DeviceID
		switch searchType {
		case "track":
			opts.URIs = append(opts.URIs, toPlay)
		case "album", "playlist", "artist":
			opts.PlaybackContext = &toPlay
		}
		//If a user tries to play a track with shuffle on,
		//they'll instead get the related tracks first
		if err := conf.Client.ShuffleOpt(ctx, false, &opts); err != nil {
			helper.LogErrorAndExit(err)
		}
		if err := conf.Client.PlayOpt(ctx, &opts); err != nil {
			helper.LogErrorAndExit(err)
		}
	},
	ValidArgs: []string{"track", "album", "playlist", "artist"},
}

func fuzzySearchResults(results spotify.SearchResult, searchType string) spotify.URI {
	var idx int
	var err error
	switch searchType {
	case "track":
		idx, err = fuzzyfinder.Find(
			results.Tracks.Tracks,
			func(i int) string {
				return fmt.Sprintf("%s - %s - %s", results.Tracks.Tracks[i].Artists[0].Name,
					results.Tracks.Tracks[i].Name,
					results.Tracks.Tracks[i].Album.Name)
			})
	case "album":
		idx, err = fuzzyfinder.Find(
			results.Albums.Albums,
			func(i int) string {
				return fmt.Sprintf("%s - %s", results.Albums.Albums[i].Name,
					results.Albums.Albums[i].Artists[0].Name)
			})
	case "playlist":
		idx, err = fuzzyfinder.Find(
			results.Playlists.Playlists,
			func(i int) string {
				return fmt.Sprintf("%s - %s", results.Playlists.Playlists[i].Name,
					results.Playlists.Playlists[i].Owner.DisplayName)
			})
	case "artist":
		idx, err = fuzzyfinder.Find(
			results.Artists.Artists,
			func(i int) string {
				return results.Artists.Artists[i].Name

			})
	}
	if err != nil {
		if err.Error() == "abort" {
			fmt.Println("Aborted search")
			os.Exit(0)
		}
		helper.LogErrorAndExit(err)
	}
	switch searchType {
	case "track":
		return results.Tracks.Tracks[idx].URI
	case "album":
		return results.Albums.Albums[idx].URI
	case "playlist":
		return results.Playlists.Playlists[idx].URI
	case "artist":
		return results.Artists.Artists[idx].URI

	}
	//The code should never get here because of our check of
	//search types earlier, this is just to make the compiler
	//happy
	return ""
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
