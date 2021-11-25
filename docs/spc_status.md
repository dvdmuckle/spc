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

