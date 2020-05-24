/*
Copyright © 2020 David Muckle <dvdmuckle@dvdmuckle>

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

	"github.com/dvdmuckle/goify/cmd/helper"
	"github.com/golang/glog"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Set device to use for all callbacks",
	Long: `Set the device to use when controlling Spotify playback.
	If this entry is empty, it will default to the currently playing device.
	You can clear the set device entry if the device is no longer active.
	This will also switch playback to the device selected if playback is active,
	and can also switch playback to the already configured device.`,
	Run: func(cmd *cobra.Command, args []string) {
		if conf.Client == (spotify.Client{}) {
			fmt.Println("Please run goify auth first to login")
			os.Exit(1)
		}
		conf.Token = *helper.RefreshToken(conf.ClientID, conf.Secret, conf.Token.RefreshToken)
		shouldClear, _ := cmd.Flags().GetBool("clear")
		shouldSwitch, _ := cmd.Flags().GetBool("noswitch")
		shouldPrint, _ := cmd.Flags().GetBool("print")
		shouldPlay, _ := cmd.Flags().GetBool("play")
		deviceToSet, _ := cmd.Flags().GetString("set")
		switch {
		case deviceToSet != "":
			setDevice(&conf, spotify.ID(deviceToSet))
			transferPlayback(&conf, shouldPlay)
		case shouldPrint:
			getDevices(&conf)
		case shouldClear:
			clearDeviceEntry(&conf)
		case !shouldSwitch:
			fuzzySwitchDevice(&conf, shouldPlay)
		default:
			transferPlayback(&conf, shouldPlay)
		}
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			glog.Fatal("Error writing config:", err)
		}
		fmt.Println("Switched to", conf.DeviceID.String())

	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	switchCmd.Flags().StringP("set", "d", "", "DeviceID to switch to")
	switchCmd.Flags().BoolP("clear", "c", false, "Clear the current device entry")
	switchCmd.Flags().BoolP("noswitch", "n", false, "Transfer playback to the currently configured device")
	switchCmd.Flags().BoolP("print", "p", false, "Only print the currently configured device")
	switchCmd.Flags().Bool("play", true, "Start playback on switch")
}

func transferPlayback(conf *helper.Config, shouldPlay bool) {
	if err := conf.Client.TransferPlayback(conf.DeviceID, shouldPlay); err != nil {
		glog.Fatal(err)
	}
}

func clearDeviceEntry(conf *helper.Config) {
	conf.DeviceID = ""
	viper.Set("device", "")
}

func getDeviceList(conf *helper.Config) []spotify.PlayerDevice {
	devices, err := conf.Client.PlayerDevices()
	if err != nil {
		glog.Fatal(err)
	}
	if len(devices) == 0 {
		fmt.Println("No devices found")
		os.Exit(0)
	}
	return devices
}

func getDevices(conf *helper.Config) {
	devices := getDeviceList(conf)
	for _, device := range devices {
		if device.ID == conf.DeviceID {
			fmt.Printf("Device configured: %s, %s\n", device.Name, device.ID.String())
			return
		}
		fmt.Println("Device configured not available, or no device is configured")
	}
}

func setDevice(conf *helper.Config, id spotify.ID) {
	conf.DeviceID = id
	viper.Set("device", conf.DeviceID.String())
}

func fuzzySwitchDevice(conf *helper.Config, shouldPlay bool) {
	devices := getDeviceList(conf)
	idx, err := fuzzyfinder.Find(
		devices,
		func(i int) string {
			switch {
			case devices[i].Active && devices[i].ID.String() == conf.DeviceID.String():
				return fmt.Sprintf("%s (currently active and configured)", devices[i].Name)
			case devices[i].Active:
				return fmt.Sprintf("%s (currently active)", devices[i].Name)
			case devices[i].ID.String() == conf.DeviceID.String():
				return fmt.Sprintf("%s (currently configured)", devices[i].Name)
			default:
				return fmt.Sprintf("%s", devices[i].Name)
			}
		})
	if err != nil {
		glog.Fatal(err)
	}
	setDevice(conf, devices[idx].ID)
	transferPlayback(conf, shouldPlay)
	return
}
