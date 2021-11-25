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
package helper

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//CreateConfig initializes a skeleton config
func CreateConfig(cfgFile string) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configPath := home + "/.config/spc"
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	if cfgFile != "" {
		if err := os.MkdirAll(configPath, 0755); err != nil {
			LogErrorAndExit("Error creating config path: ", err)
		}
		cfgFile = fmt.Sprintf(configPath + "/config.yaml")
	}
	viper.SetDefault("spotifyclientid", "Your Spotify ClientID")
	viper.SetDefault("spotifysecret", "Your Spotify Client Secret base64 encoded")
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			LogErrorAndExit("Error writing config file:", err)
		}
		fmt.Printf("Config file created at ~/.config/spc/config.yaml\n")
	}
}

//SetupConfig sets up the path and type of the config for Viper
//and also returns the full path to the config
func SetupConfig() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configPath := home + "/.config/spc"
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return configPath + "/config.yaml"
}
