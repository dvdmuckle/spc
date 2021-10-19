## spc seek

Seek to a specific position in the currently playing song from Spotify

### Synopsis

Seek to a specific position in the currently playing song from Spotify. This command requires
exactly one argument, either a number between 0 and the length of the song in seconds, or a timestamp in
the form of minutes:seconds.

```
spc seek [flags]
```

### Options

```
  -h, --help   help for seek
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.config/spc/config.yaml)
  -v, --verbose         verbose error logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

