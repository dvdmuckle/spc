## spc save-weekly

Saves the current Spotify Discover Weekly playlist

### Synopsis

Saves the current Spotify Discover Weekly playlist.
Note this cannot bring back old Spotify Discover Weekly playlists, it can
only save the current playlist

```
spc save-weekly [flags]
```

### Options

```
  -h, --help          help for save-weekly
  -n, --name string   Custom name for the save playlist
  -p, --public        Whether to make the new playlist public
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --config string                    Config file (default is $HOME/.config/spc/config.yaml)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

