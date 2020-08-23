package helper

import (
	"fmt"
	"os"

	"github.com/golang/glog"
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
			glog.Fatal("Error creating config path: ", err)
		}
		cfgFile = fmt.Sprintf(configPath + "/config.yaml")
	}
	viper.SetDefault("spotifyclientid", "Your Spotify ClientID")
	viper.SetDefault("spotifysecret", "Your Spotify Client Secret base64 encoded")
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			glog.Fatal("Error writing config file:", err)
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
	return configPath + "config.yaml"
}
