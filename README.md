# goify [Tentative Name]

![Go](https://github.com/dvdmuckle/goify/workflows/Go/badge.svg?branch=master)

A lightweight mutliplatform CLI for Spotify

## Setup

To set up the app, run `goify` to generate a skeleton config file at `~/.config/goify/config.yaml`. Next, head to <http://developer.spotify.com/> to create a new Spotify app. Paste the ClientID and ClientSecret in the config as noted. You can now run `goify auth` to start the OAuth2 flow, which will have you grant the Spotify app you created, and thus goify, the correct API permissions.

## Running

Because this app is in constant development, please refer to `goify help` for what functionality is available in the app.

## Goals

The goal of this project is to present a simple, lightweight command line interface for Spotify, inspired by [spotify-tui](https://github.com/Rigellute/spotify-tui). Support for play, pause, volume, and a simple search are all that are considered right now. The roadmap may change in the future.

## Roadmap

| Feature | Implemented yet? | Essential? |
|---------|------------------|------------|
| Auth | Yes | Yes |
| Switch Device | Yes | No |
| Play | No | Yes |
| Pause | No | Yes |
| Toggle playback | No | No |
| Search | No | Yes |
| Volume | No | Yes |
| Status | No | Yes |

## Contributing

For feature requests, feel free to create an issue or creating a PR changing the above roadmap.
