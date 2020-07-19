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
	"strings"

	"github.com/dvdmuckle/goify/cmd/helper"
	"github.com/golang/glog"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		searchTerm := strings.Join(args, " ")
		searchResults, err := conf.Client.Search(searchTerm, spotify.SearchTypeTrack)
		if err != nil {
			glog.Fatal(err)
		}
		toPlay := fuzzySearchResults(*searchResults)
		conf.Client.QueueSong(toPlay)
	},
}

func fuzzySearchResults(results spotify.SearchResult) spotify.ID {
	idx, err := fuzzyfinder.Find(
		results.Tracks.Tracks,
		func(i int) string {
			return fmt.Sprintf("%s - %s", results.Tracks.Tracks[i].Artists[0].Name, results.Tracks.Tracks[i].Name)
		})

	if err != nil {
		if err.Error() == "abort" {
			fmt.Println("Aborted switch")
			os.Exit(0)
		}
		glog.Fatal(err)
	}
	fmt.Println(results.Tracks.Tracks[idx].Name)
	return results.Tracks.Tracks[idx].ID
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
