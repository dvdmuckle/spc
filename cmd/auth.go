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
	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticates with Spotify",
	Long: `Authenticates with Spotify by printout out a login link, which will then save your access token to the config file.
Use this command after the initial login to refresh your access token`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.Auth(cmd, cfgFile, &conf)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().BoolP("refresh", "r", false, "Force refreshing the token")
}
