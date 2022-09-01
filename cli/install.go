package cli

import "github.com/gookit/gcli/v3"

// InstallCmd define
var InstallCmd = &gcli.Command{
	Name:    "install",
	Desc:    "install current Go version",
	Aliases: []string{"i", "ins"},
	Config: func(c *gcli.Command) {

	},
	Func: func(c *gcli.Command, args []string) error {
		return nil
	},
}
