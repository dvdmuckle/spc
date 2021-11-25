## spc auth

Authenticates with Spotify

### Synopsis

Authenticates with Spotify by printout out a login link, which will then save your access token to the config file.
Use this command after the initial login to refresh your access token

```
spc auth [flags]
```

### Options

```
  -h, --help      help for auth
  -r, --refresh   Force refreshing the token
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.config/spc/config.yaml)
  -v, --verbose         verbose error logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

