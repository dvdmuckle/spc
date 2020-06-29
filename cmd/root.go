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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dvdmuckle/goify/cmd/helper"
	"github.com/golang/glog"
	"github.com/zmb3/spotify"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//Config type stores constantly retrieved things from the config file

var conf helper.Config

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goify",
	Short: "Command line tool to control Spotify",
	Long: `Goify is a simple command line tool to control Spotify using the Spotify API
to allow for cross platform use. It is designed to be simple and is limited in
scope, and is best when paired with another more complicated tool.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	createConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.config/goify/config.yaml)")
}

// createConfig creates the config file at ~/.config/goify/config.yaml if it does not exist
func createConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configPath := home + "/.config/goify"
	if err := os.MkdirAll(configPath, 0755); err != nil {
		glog.Fatal("Error creating config path: ", err)
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	cfgFile = fmt.Sprintf(configPath + "/config.yaml")
	viper.SetDefault("spotifyclientid", "Your Spotify ClientID")
	viper.SetDefault("spotifysecret", "Your Spotify Client Secret base64 encoded")
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			glog.Fatal("Error writing config file:", err)
		}
		fmt.Printf("Config file created at ~/.config/goify/config.yaml\n\n")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		createConfig()
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		glog.Fatal("Error reading config file:", err)
	}
	viper.AutomaticEnv() // read in environment variables that match
	conf.ClientID = viper.GetString("spotifyclientid")
	if secret, err := base64.StdEncoding.DecodeString(viper.GetString("spotifysecret")); err != nil && len(secret) != 0 {
		glog.Fatal("Error decoding Spotify Client Secret, is it valid and base64 encoded? Error: ", err)
	} else {
		conf.Secret = strings.TrimSpace(string(secret))
	}
	if viper.GetString("auth") != "" {
		if err := json.Unmarshal([]byte(viper.GetString("auth")), &conf.Token); err != nil {
			glog.Fatal(err)
		}
	}
	conf.DeviceID = spotify.ID(viper.GetString("device"))
}
