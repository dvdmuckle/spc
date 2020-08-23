/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Initializes a config file",
	Long: `Creates a config file in ~/.config/spc/config.yaml.
	The config file will then have to be manually adjusted to add the
	Spotify ClientID and the Spotify Client Secret as noted in the 
	config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.CreateConfig(cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
