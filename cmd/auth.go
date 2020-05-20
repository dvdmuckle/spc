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
	"net/http"
	"os"

	"github.com/dvdmuckle/goify/cmd/helper"
	"github.com/golang/glog"
	spotifyAuth "github.com/markbates/goth/providers/spotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8888/callback"

var (
	authenticator = spotify.NewAuthenticator(redirectURI)
	ch            = make(chan *spotify.Client)
	clientID      string
	secret        string
	state         = "ringdingthing"
)

func auth() {
	clientID = viper.GetString("spotifyclientid")
	secret = viper.GetString("spotifysecret")
	if clientID == "" || secret == "" {
		fmt.Println("Please configure your Spotify client ID and secret in the config file at ~/.config/goify/config.yaml")
		os.Exit(1)
	}
	provider := spotifyAuth.New(clientID, secret, redirectURI)
	if viper.GetString("auth.token") != "" && provider.RefreshTokenAvailable() {
		viper.Set("auth.token", helper.RefreshToken(provider, viper.GetString("auth.token")))
	} else {
		fmt.Println("Getting token...")
		authenticator.SetAuthInfo(clientID, secret)
		http.HandleFunc("/callback", completeAuth)
		go http.ListenAndServe(":8888", nil)
		url := authenticator.AuthURL(state)
		fmt.Println("Please log in to Spotify by clicking the following link:", url)
		//wait for auth to finish
		client := <-ch

		user, err := client.CurrentUser()
		if err != nil {
			glog.Fatal(err)
		}
		fmt.Println("Login successful as ", user.ID)
	}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := authenticator.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		glog.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		glog.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := authenticator.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		auth()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
