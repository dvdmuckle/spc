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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	spotifyAuth "github.com/markbates/goth/providers/spotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:8888/callback"

var (
	authenticator = spotify.NewAuthenticator(redirectURI, spotify.ScopeStreaming, spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadPlaybackState, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
	ch            = make(chan *spotify.Client)
	clientID      string
	secret        string
	state         = "ringdingthing"
)

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

//Auth authenticates with Spotify and refreshes the token
func Auth(cmd *cobra.Command, viper *viper.Viper, cfgFile string, conf *Config) {
	clientID = conf.ClientID
	secret = conf.Secret
	if clientID == "" || secret == "" {
		fmt.Println("Please configure your Spotify client ID and secret in the config file at ~/.config/spc/config.yaml")
		os.Exit(1)
	}

	shouldRefresh, err := cmd.Flags().GetBool("refresh")
	if err != nil {
		glog.Fatal(err)
	}
	if len(viper.GetString("auth")) != 0 && shouldRefresh {
		fmt.Println("Refreshing token...")
		newToken := RefreshToken(clientID, secret, conf.Token.RefreshToken)
		conf.Token = *newToken
		marshalToken, err := json.Marshal(conf.Token)
		if err != nil {
			glog.Fatal(err)
		}
		viper.Set("auth", string(marshalToken))
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			glog.Fatal("Error writing config:", err)
		}
	} else {
		fmt.Println("Getting token...")
		authenticator.SetAuthInfo(clientID, secret)
		http.HandleFunc("/callback", completeAuth)
		go http.ListenAndServe(":8888", nil)
		url := authenticator.AuthURL(state)
		fmt.Println("Please log in to Spotify by clicking the following link, or copying it to a web browser:", url)
		//wait for auth to finish
		client := <-ch

		user, err := client.CurrentUser()
		if err != nil {
			glog.Fatal(err)
		}
		token, err := client.Token()
		if err != nil {
			glog.Fatal(err)
		}
		conf.Token = *token
		marshalToken, err := json.Marshal(conf.Token)
		if err != nil {
			glog.Fatal(err)
		}
		viper.Set("auth", string(marshalToken))
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			glog.Fatal("Error writing config:", err)
		}
		fmt.Println("Login successful as", user.ID)
	}
}

//RefreshToken refreshes the auth token from Spotify
//TODO: #4 Replace implementation with vanilla oauth2 use
func RefreshToken(client string, secret string, refreshToken string) *oauth2.Token {
	provider := spotifyAuth.New(client, secret, redirectURI)
	if refreshToken != "" && provider.RefreshTokenAvailable() {
		token, err := provider.RefreshToken(refreshToken)
		if err != nil {
			glog.Fatal(err)
		}
		return token
	}
	glog.Fatal("Cannot refresh token, token is empty")
	return nil
}

//SetClient sets the Client field of Config struct to a valid Spotify client
//The Token field in the Config struct must be set
func SetClient(conf *Config, cfgFile string) {
	if conf.Token == (oauth2.Token{}) {
		fmt.Println("Please run spc auth first to login")
		os.Exit(1)
	}
	if time.Now().After(conf.Token.Expiry) {
		conf.Token = *RefreshToken(conf.ClientID, conf.Secret, conf.Token.RefreshToken)
		marshalToken, err := json.Marshal(conf.Token)
		if err != nil {
			glog.Fatal(err)
		}
		viper.Set("auth", string(marshalToken))
		if err := viper.WriteConfigAs(cfgFile); err != nil {
			glog.Fatal("Error writing config:", err)
		}
	}
	conf.Client = spotify.NewAuthenticator(redirectURI).NewClient(&conf.Token)
}
