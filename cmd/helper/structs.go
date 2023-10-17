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
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

// Config stores constantly accessed variables in memory
type Config struct {
	ClientID string
	Secret   string
	Token    oauth2.Token
	Client   *spotify.Client
	DeviceID spotify.ID
}
