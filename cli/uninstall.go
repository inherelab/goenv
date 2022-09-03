package cli

import (
	"github.com/gookit/gcli/v3"
)

// UninstallCmd define
var UninstallCmd = &gcli.Command{
	Name:    "uninstall",
	Desc:    "uninstall input Go version",
	Aliases: []string{"un", "uni", "rm"},
	Config: func(c *gcli.Command) {
		c.AddArg("version", "want uninstalled go version", true)
	},
	Func: func(c *gcli.Command, args []string) error {
		ver := c.Arg("version").String()
		ver, err := checkFormatVersion(ver)
		if err != nil {
			return err
		}

		adaptor, err := makeAdaptor()
		if err != nil {
			return err
		}

		return adaptor.Uninstall(ver)
	},
}
