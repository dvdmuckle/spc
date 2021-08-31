# spc

![Go](https://github.com/dvdmuckle/spc/workflows/Go/badge.svg?branch=master)

A lightweight mutliplatform CLI for Spotify

## Installation

### macOS

```bash
brew install dvdmuckle/tap/spc
```

### Ubuntu Linux

```bash
sudo add-apt-repository ppa:dvdmuckle/spc
sudo apt-get update
sudo apt-get install spc
```

### Fedora Linux

```bash
sudo dnf copr enable dvdmuckle/spc
sudo dnf install spc
```

### Other Platforms

Download one of the releases and unarchive the `spc` binary somewhere in your `PATH`. Alternatively, download and install with:
``
go get -u github.com/dvdmuckle/spc
``
Make sure `$GOPATH/bin` is in your `PATH` for this to work.

## Setup

To set up the app, run `spc config` to generate a skeleton config file at `~/.config/spc/config.yaml` or `.config/spc/config.yaml` in your user directory on Windows. Next, head to <http://developer.spotify.com/dashboard> to create a new Spotify app. Make sure to set a callback URL for `http://localhost:8888/callback`. Paste the ClientID and ClientSecret in the newly created config file as noted. Make sure the ClientSecret is base64 encoded. You can now run `spc auth` to start the OAuth2 flow, which will have you grant the Spotify app you created, and thus spc, the correct API permissions.

**This app requires a Spotify Premium account for any commands involving playback.**

## Running

Check out either `spc help` or the [docs pages](docs/spc.md) for help on how to use spc.

## Goals

The goal of this project is to present a simple, lightweight command line interface for Spotify, inspired by [spotify-tui](https://github.com/Rigellute/spotify-tui). Support for play, pause, volume, and a simple search are all that are considered right now. More complicated tasks like managing playlists are not considered at this time.

## Contributing

For feature requests, feel free to create an issue or submit a PR with your changes.
