## spc docs

Generates docs fpr spc

### Synopsis

Generates docs for spc.
This command is mostly used for automation purposes, but can be used to generate
either man page or markdown documentation. The first argument is which
kind of documentation to generate, either man or markdown. The second is the path for the
generated docs. If the path does not exist, it will be created.

```
spc docs [flags]
```

### Options

```
      --gen-tags   Add autogentags to generated docs
  -h, --help       help for docs
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.config/spc/config.yaml)
  -v, --verbose         verbose error logging
```

### SEE ALSO

* [spc](spc.md)	 - Command line tool to control Spotify

