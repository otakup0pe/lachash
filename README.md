[![Build Status](https://travis-ci.org/otakup0pe/lachash.svg?branch=master)](https://travis-ci.org/otakup0pe/lachash)

# Lachash

This tool provides a means of securely exchanging arbitrary data with others via a Hashicorp Vault deployment. It is designed with some basic old-school UNIX in mind, and is currently meant for sharing small snippets of data as opposed to giant files.

The mechanics are fairly simple. When writing the data to Vault a context-specific token is created. This is used to store the data in a cubbyhole, and subsequently retrieve the data. The receiving party doesn't actually even need a Vault account as the provided token is the accesor.

# Install

You can download premade binaries on the [releases](https://github.com/otakup0pe/lachash/releases/latest) page. These zip files contain the precompiled Go binary.

If you wish to compile it yourself, there is a [Makefile](https://github.com/otakup0pe/lachash/blob/master/Makefile) which can be invoked with no target to run tests and build the various binaries.

# Commands

## `stash`

The stash command is used to write something into Vault. If given no arguments, it will attempt to read from stdin. You can also read from a file and change a variety of options such as TTL and how many times the data may be read.

By default, you will get a normal Vault token back. You may also make use of a "short code" by passing the `-short-code` option.

```
lachash stash -input /tmp/cornbread_recipe.txt
ba3591b4-08ca-0dfd-9861-16d4957c2060
echo "just kidding" | lachash stash -short-code
/6a73DcaGRp5dzLQ7qdnkg
```

## `pop`

The pop command is used to pull something from Vault. Unless an output file is specified, the data will be written to stdout. You should probably pass in either the `-token` or `-short-code` options. Failure to do so will result in `lachash` trying to use your normal environmental Vault credentials, which is generally not desired outside of automation contexts.

```
lachash pop -output /tmp/cornbread_recipe.txt -token ba3591b4-08ca-0dfd-9861-16d4957c2060
lachash pop -short-code /6a73DcaGRp5dzLQ7qdnkg
just kidding
```

# License

[Apache 2](https://github.com/otakup0pe/lachash/blob/master/LICENSE)

# Author

The lachash tool was created by [Jonathan Freedman](http://jonathanfreedman.bio/).
