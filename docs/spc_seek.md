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

