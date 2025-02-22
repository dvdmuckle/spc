## spc search

Search for and play a track, album, playlist, or artist

### Synopsis

Search takes two arguments: the search type, and the query.
Search type can be an album, a track, a playlist, or an artist, with the rest of the arguments
making up the search query. For example:

	spc search album moving pictures
	spc search track tom sawyer
	spc search playlist prog monsters
	spc search artist rush

If a track is queried for, additional similar songs will be queued up.

More advanced options are available for the search query. For this,
please see https://pkg.go.dev/github.com/zmb3/spotify?tab=doc#Client.Search

```
spc search [flags]
```

### Options

```
  -h, --help   help for search
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.config/spc/config.yaml)
  -v, --verbose         verbose error logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

