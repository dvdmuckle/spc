package helper

import "golang.org/x/oauth2"

//Config stores constantly accessed variables in memory
type Config struct {
	ClientID string
	Secret   string
	Token    oauth2.Token
}
