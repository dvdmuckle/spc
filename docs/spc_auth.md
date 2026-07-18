## spc auth

Authenticates with Spotify

### Synopsis

Authenticates with Spotify by printing a login link, which will then save your access token to the config file.
Use this command after the initial login to refresh your access token.

Before running this command, make sure your Spotify app's callback URL is set to `http://127.0.0.1:8888/callback` in the Spotify developer dashboard.

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

