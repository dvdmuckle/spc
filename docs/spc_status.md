## spc status

Get the currently playing song from Spotify

### Synopsis

Get the currently playing song from Spotify

A format string can be passed with --format or -f to specify what
the status printout should look like. The following fields are available:

%a - Artist
%t - Track
%b - Album
%f - Playing
%e ️- ▶ or ⏸️
%s - Play progress

If a song has multiple artists, you can specify the upper limit of artists
to display with %Xa, where X is the number of artists to print, separated
by commas.

If there is no currently playing song on Spotify, regardless of format argument
the command will return an empty string. This may happen if Spotify is paused
for an extended period of time

```
spc status [flags]
```

### Options

```
  -f, --format string   Format string for formatting the status
  -h, --help            help for status
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.config/spc/config.yaml)
  -v, --verbose         verbose error logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

