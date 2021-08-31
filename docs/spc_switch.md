## spc switch

Set device to use for all callbacks

### Synopsis

Set the device to use when controlling Spotify playback.
If this entry is empty, it will default to the currently playing device.
You can clear the set device entry if the device is no longer active.
This will also switch playback to the device selected if playback is active,
and can also switch playback to the already configured device.

```
spc switch [flags]
```

### Options

```
  -c, --clear           Clear the current device entry
  -h, --help            help for switch
  -p, --play            Start playback on switch
      --print           Only print the currently configured device
  -d, --set string      DeviceID to switch to
  -t, --transfer-only   Transfer playback to the currently configured device
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

###### Auto generated by spf13/cobra on 31-Aug-2021