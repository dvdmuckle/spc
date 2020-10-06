# spc

![Go](https://github.com/dvdmuckle/spc/workflows/Go/badge.svg?branch=master)

A lightweight mutliplatform CLI for Spotify

## Installation

### macOS

```bash
brew install dvdmuckle/tap/spc
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

To set up the app, run `spc config` to generate a skeleton config file at `~/.config/spc/config.yaml`. Next, head to <http://developer.spotify.com/> to create a new Spotify app. Make sure to set a callback URL for `http://localhost:8888/callback`. Paste the ClientID and ClientSecret in the config as noted. Make sure the ClientSecret is base64 encoded. You can now run `spc auth` to start the OAuth2 flow, which will have you grant the Spotify app you created, and thus spc, the correct API permissions.

**This app requires a Spotify Premium account in order to interact with the API.**

## Running

Because this app is in constant development, please refer to `spc help` for what functionality is available in the app.

## Goals

The goal of this project is to present a simple, lightweight command line interface for Spotify, inspired by [spotify-tui](https://github.com/Rigellute/spotify-tui). Support for play, pause, volume, and a simple search are all that are considered right now. The roadmap may change in the future.

## Roadmap

| Feature | Implemented yet? | Essential? |
|---------|------------------|------------|
| Auth | Yes | Yes |
| Switch Device | Yes | No |
| Play | Yes | Yes |
| Pause | Yes | Yes |
| Toggle playback | No | No |
| Search | Yes | Yes |
| Volume | Yes | Yes |
| Status | Yes | Yes |
| Skip Track | Yes | Yes |
| Previous Track | Yes | Yes |
| Config | Yes | No |

## Contributing

For feature requests, feel free to create an issue or creating a PR changing the above roadmap.
