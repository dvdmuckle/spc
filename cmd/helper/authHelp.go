package helper

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang/glog"
	spotifyAuth "github.com/markbates/goth/providers/spotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8888/callback"

var (
	authenticator = spotify.NewAuthenticator(redirectURI, spotify.ScopeStreaming, spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadPlaybackState)
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
func Auth(cmd *cobra.Command, viper *viper.Viper, cfgFile string) {
	clientID = viper.GetString("spotifyclientid")
	secret = viper.GetString("spotifysecret")
	if clientID == "" || secret == "" {
		fmt.Println("Please configure your Spotify client ID and secret in the config file at ~/.config/goify/config.yaml")
		os.Exit(1)
	}

	provider := spotifyAuth.New(clientID, secret, redirectURI)
	shouldRefresh, err := cmd.Flags().GetBool("refresh")
	if err != nil {
		glog.Fatal(err)
	}
	if viper.GetString("auth.accesstoken") != "" && shouldRefresh {
		fmt.Println("Refreshing token...")
		viper.Set("auth.accesstoken", RefreshToken(provider, viper.GetString("auth.refreshtoken")))
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
		token, err := client.Token()
		if err != nil {
			glog.Fatal(err)
		}
		viper.Set("auth", token)
		viper.WriteConfigAs(cfgFile)
		fmt.Println("Login successful as", user.ID)
	}
}

//RefreshToken refreshes the auth token from Spotify
func RefreshToken(provider *spotifyAuth.Provider, tokenToRefresh string) string {
	if tokenToRefresh != "" && provider.RefreshTokenAvailable() {
		token, err := provider.RefreshToken(tokenToRefresh)
		if err != nil {
			glog.Fatal(err)
		}
		return token.AccessToken
	}
	glog.Fatal("Cannot refresh token, token is empty")
	return ""
}
