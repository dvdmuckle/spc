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
      --config          Switch configured device but do not transfer playback
  -h, --help            help for switch
  -p, --play            Start playback on switch
      --print           Only print the currently configured device
  -d, --set string      DeviceID to switch to
  -t, --transfer-only   Transfer playback to the currently configured device
```

### Options inherited from parent commands

```
  -v, --verbose   verbose error logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

