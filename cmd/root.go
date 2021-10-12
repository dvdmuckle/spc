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
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/zmb3/spotify"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

//Config type stores constantly retrieved things from the config file

var conf helper.Config

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spc",
	Short: "Command line tool to control Spotify",
	Long: `Spc is a simple command line tool to control Spotify using the Spotify API
to allow for cross platform use. It is designed to be simple and is limited in
scope, and is best when paired with another more complicated tool.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//So glog doesn't tell us we're logging before parsing flags
		//This is entirely bogus since it's parsing an empty string array
		//Plug cobra handles all our flags anyways
		flag.CommandLine.Parse([]string{})
	},
	DisableAutoGenTag: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.config/spc/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(helper.GetVerboseErrLogAddr(), "verbose", "v", false, "verbose error logging")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		_ = helper.SetupConfig()
		viper.SetConfigFile(cfgFile)
	} else {
		cfgFile = helper.SetupConfig()
	}
	// If a config file is found, read it in.
	if _, err := os.Stat(cfgFile); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			helper.LogErrorAndExit("Error reading config file: ", err)
		}
	}
	viper.AutomaticEnv() // read in environment variables that match
	conf.ClientID = viper.GetString("spotifyclientid")
	if secret, err := base64.StdEncoding.DecodeString(viper.GetString("spotifysecret")); err != nil && len(secret) != 0 {
		helper.LogErrorAndExit("Error decoding Spotify Client Secret, is it valid and base64 encoded? Error: ", err)
	} else {
		conf.Secret = strings.TrimSpace(string(secret))
	}
	conf.DeviceID = spotify.ID(viper.GetString("device"))
}
