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
	"os"
	"reflect"
	"strings"

	goflag "flag"

	"github.com/dvdmuckle/spc/cmd/helper"
	"github.com/zmb3/spotify"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

//Config type stores constantly retrieved things from the config file

var conf helper.Config

var cfgFile string

var verboseErrLog bool

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
		helper.LogErrorAndExit(false, err)
	}
}

type flagValueWrapper struct {
	inner    goflag.Value
	flagType string
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.config/spc/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verboseErrLog, "verbose", "v", false, "verbose error logging")
	//fmt.Println(flag.CommandLine.Lookup("v").Value)
	//fmt.Println(flag.CommandLine.Lookup("logtostderr").Value)
	flag.CommandLine.Lookup("v")
	//rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	vFlag := flag.CommandLine.Lookup("v")

	//fmt.Println(vFlag.Value)
	pv := &flagValueWrapper{
		inner: vFlag.Value,
	}
	//fmt.Println(pv.flagType)
	t := reflect.TypeOf(vFlag.Value)
	if t.Kind() == reflect.Interface || t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	//fmt.Println(t)
	//fmt.Println(t.Name())
	pv.flagType = strings.TrimSuffix(t.Name(), "Value")

	//fmt.Println(pv.flagType)
	//fmt.Println(rootCmd.PersistentFlags().Lookup("logtostderr").Value.Type())
	//fmt.Println(rootCmd.PersistentFlags().Lookup("v").Value.Type())
	//vFlag := rootCmd.PersistentFlags().Lookup("v")
	//vFlagString := vFlag.Value.String()
	//vFlagSet := vFlag.Value.Set()
	//vFlag.Value = pflag.Value{vFlagString, false, "count"}
	//fmt.Println(flag.CommandLine.Lookup("v"))
	//fmt.Println(rootCmd.PersistentFlags().Lookup("config"))

	//verbose = flag.CommandLine.Lookup("v").Value
	//fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	//flag.PrintDefaults()
	//fmt.Println(flag.CommandLine.Lookup("v"))
	//fmt.Println(rootCmd.Flag("v"))
	//fmt.Println(rootCmd.Flag("config"))

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
			helper.LogErrorAndExit(verboseErrLog, "Error reading config file: ", err)
		}
	}
	viper.AutomaticEnv() // read in environment variables that match
	conf.ClientID = viper.GetString("spotifyclientid")
	if secret, err := base64.StdEncoding.DecodeString(viper.GetString("spotifysecret")); err != nil && len(secret) != 0 {
		helper.LogErrorAndExit(verboseErrLog, "Error decoding Spotify Client Secret, is it valid and base64 encoded? Error: ", err)
	} else {
		conf.Secret = strings.TrimSpace(string(secret))
	}
	conf.DeviceID = spotify.ID(viper.GetString("device"))
}
