package helper

import (
	"fmt"

	"github.com/golang/glog"
	spotifyAuth "github.com/markbates/goth/providers/spotify"
)

//RefreshToken refreshes the auth token from Spotify
func RefreshToken(provider *spotifyAuth.Provider, tokenToRefresh string) string {
	if tokenToRefresh != "" && provider.RefreshTokenAvailable() {
		fmt.Println("Refreshing token...")
		token, err := provider.RefreshToken(tokenToRefresh)
		if err != nil {
			glog.Fatal(err)
		}
		return token.AccessToken
	}
	glog.Fatal("Cannot refresh token, token is empty")
	return ""
}
