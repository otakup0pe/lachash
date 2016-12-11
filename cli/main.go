package cli

import (
	"github.com/mitchellh/cli"
)

func CommandFactory() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"stash": func() (cli.Command, error) {
			return &Stash{}, nil
		},
		"pop": func() (cli.Command, error) {
			return &Pop{}, nil
		},
	}
}

