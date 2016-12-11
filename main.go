package main

import (
	"os"
	"github.com/otakup0pe/lachash/helpers"
	cli_mod "github.com/mitchellh/cli"
	"github.com/otakup0pe/lachash/cli"
)
var lachash_version string = "unknown"

func main() {
	helpers.Init()
	c := cli_mod.NewCLI("lachash", Version())
	c.Args = os.Args[1:]
	c.Commands = cli.CommandFactory()
	rc, err := c.Run()
	if err != nil {
		helpers.Log(err.Error())
	}

	os.Exit(rc)
}

func Version() string {
	return lachash_version
}
