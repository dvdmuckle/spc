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
package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:8888/callback"

var (
	authenticator *spotifyauth.Authenticator
	ch       = make(chan *spotify.Client)
	clientID string
	secret   string
	state    = "ringdingthing"
)

func initAuthenticator(clientID string, secret string) {
    authenticator = spotifyauth.New(
        spotifyauth.WithRedirectURL(redirectURI),
        spotifyauth.WithClientID(clientID),
        spotifyauth.WithClientSecret(secret),
        spotifyauth.WithScopes(
            spotifyauth.ScopeStreaming,
            spotifyauth.ScopeUserModifyPlaybackState,
            spotifyauth.ScopeUserReadPlaybackState,
            spotifyauth.ScopePlaylistModifyPrivate,
            spotifyauth.ScopePlaylistModifyPublic,
        ),
    )
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := authenticator.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		LogErrorAndExit(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		LogErrorAndExit(fmt.Sprintf("State mismatch: %s != %s\n", st, state))
	}
	// use the token to get an authenticated client
	client := spotify.New(authenticator.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

// Auth authenticates with Spotify and refreshes the token
func Auth(cmd *cobra.Command, cfgFile string, conf *Config) {
	clientID = conf.ClientID
	secret = conf.Secret
	initAuthenticator(clientID, secret)
	curUser, err := user.Current()
	if err != nil {
		LogErrorAndExit(err)
	}
	if clientID == "" || secret == "" {
		fmt.Println("Please configure your Spotify client ID and secret in the config file at ~/.config/spc/config.yaml")
		os.Exit(1)
	}

	shouldRefresh, err := cmd.Flags().GetBool("refresh")
	if err != nil {
		LogErrorAndExit(err)
	}
	if key, err := keyring.Get("spc", curUser.Username); err == nil && key != "" && shouldRefresh {
		fmt.Println("Refreshing token...")
		if err := json.Unmarshal([]byte(key), &conf.Token); err != nil {
			LogErrorAndExit(err)
		}
		newToken := RefreshToken(clientID, secret, conf.Token.RefreshToken)
		conf.Token = *newToken
		marshalToken, err := json.Marshal(conf.Token)
		if err != nil {
			LogErrorAndExit(err)
		}
		if err := keyring.Set("spc", curUser.Username, string(marshalToken)); err != nil {
			LogErrorAndExit("Error saving token to keyring", err)
		}
	} else {
		fmt.Println("Getting token...")
		http.HandleFunc("/callback", completeAuth)
		go http.ListenAndServe(":8888", nil)
		url := authenticator.AuthURL(state)
		fmt.Println("Please log in to Spotify by clicking the following link, or copying it to a web browser:", url)
		//wait for auth to finish
		client := <-ch
		if client == nil {
			fmt.Println("Client is not initialized")
			os.Exit(1)
		}

		user, err := client.CurrentUser(context.Background())
		if err != nil {
			LogErrorAndExit(err)
		}
		token, err := client.Token()
		if err != nil {
			LogErrorAndExit(err)
		}
		conf.Token = *token
		marshalToken, err := json.Marshal(conf.Token)
		if err != nil {
			LogErrorAndExit(err)
		}
		if err := keyring.Set("spc", curUser.Username, string(marshalToken)); err != nil {
			LogErrorAndExit("Error saving token to keyring", err)
		}
		fmt.Println("Login successful as", user.ID)
	}
}

// RefreshToken refreshes the auth token from Spotify
func RefreshToken(client string, secret string, refreshToken string) *oauth2.Token {
	var token *oauth2.Token = &oauth2.Token{}

	if refreshToken != "" {
		const grantType string = "refresh_token"
		const tokenURL string = spotifyauth.TokenURL
		const contentType string = "application/x-www-form-urlencoded"
		form := url.Values{}
		form.Add("grant_type", grantType)
		form.Add("refresh_token", refreshToken)
		form.Add("client_id", client)
		form.Add("client_secret", secret)

		resp, err := http.Post(tokenURL, contentType, strings.NewReader(form.Encode()))
		if err != nil {
			LogErrorAndExit(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			LogErrorAndExit(err)
		}
		json.Unmarshal(body, token)
		//Looks like the normal token refresh from Spotify doesn't also return the refreshToken
		token.RefreshToken = refreshToken
		// TODO Don't hardcode the expiry duration
		token.Expiry = time.Now().Add(time.Second * 3600)
		return token
	}
	LogErrorAndExit("Cannot refresh token, token is empty")
	return nil
}

// SetClient sets the Client field of Config struct to a valid Spotify client
// The Token field in the Config struct must be set
func SetClient(conf *Config) {
	curUser, err := user.Current()
	if err != nil {
		LogErrorAndExit(err)
	}
	if key, err := keyring.Get("spc", curUser.Username); err == nil && key != "" {
		if err := json.Unmarshal([]byte(key), &conf.Token); err != nil {
			LogErrorAndExit(err)
		}
	} else {
		fmt.Println("Please run spc auth first to login")
		os.Exit(1)
	}
	//I'm 99% certain this isn't a case we can run into, but still...
	if conf.Token == (oauth2.Token{}) {
		fmt.Println("Please run spc auth first to login")
		os.Exit(1)
	}
	if time.Now().After(conf.Token.Expiry) {
		conf.Token = *RefreshToken(conf.ClientID, conf.Secret, conf.Token.RefreshToken)
		marshalToken, err := json.Marshal(conf.Token)
		if err != nil {
			LogErrorAndExit(err)
		}
		curUser, err := user.Current()
		if err != nil {
			LogErrorAndExit(err)
		}
		if err := keyring.Set("spc", curUser.Username, string(marshalToken)); err != nil {
			LogErrorAndExit("Error saving token to keyring", err)
		}
	}
	httpClient := spotifyauth.New().Client(context.Background(), &conf.Token)
	conf.Client = spotify.New(httpClient)
	if conf.Client == nil {
		fmt.Println("Client is not initialized")
		os.Exit(1)
	}
}
