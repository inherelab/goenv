package cli

import (
	"github.com/gookit/gcli/v3"
	"github.com/inherelab/goenv"
	"github.com/inherelab/goenv/internal"
)

// UninstallCmd define
var UninstallCmd = &gcli.Command{
	Name:    "uninstall",
	Desc:    "uninstall input Go version",
	Aliases: []string{"un", "uni"},
	Config: func(c *gcli.Command) {
		c.AddArg("version", "want uninstalled go version", true)
	},
	Func: func(c *gcli.Command, args []string) error {
		ver := c.Arg("version").String()

		adaptor, err := internal.NewEnvManager(goenv.Cfg.Mode).CreateAdaptor()
		if err != nil {
			return err
		}

		return adaptor.Uninstall(ver)
	},
}
