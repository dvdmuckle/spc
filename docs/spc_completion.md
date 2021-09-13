## spc completion

Generate shell autocompletion

### Synopsis

Generates shell autocompletion

The following shells can have autocompletion generated:

bash
zsh
fish
powershell

For fish, the flag --fish-description can be toggled to includes descriptions in the autocomplete

```
spc completion [flags]
```

### Options

```
      --fish-description   Whether to include description for fish autocompletion
  -h, --help               help for completion
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

