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
	"strconv"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
)

// volumeCmd represents the volume command
var volumeCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"vol"},
	Args:    cobra.ExactArgs(1),
	Short:   "Set volume for Spotify",
	Long: `Sets the volume for Spotify. This command requires exactly
one argument, a number between 0 and 100 to set the volume to.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.SetClient(&conf)
		vol, err := strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			fmt.Println("Volume is not a number")
			os.Exit(1)
		}
		if vol > 100 {
			fmt.Println("Volume cannot be higher than 100")
			os.Exit(1)
		} else if vol < 0 {
			//By virtue of how cobra processes args, where "-" denotes a flag, we should never get here
			//Still, just in case...
			fmt.Println("Volume cannot be less than 0")
			os.Exit(1)
		}
		conf.Client.VolumeOpt(context.Background(), int(vol), &spotify.PlayOptions{DeviceID: &conf.DeviceID})
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
}
