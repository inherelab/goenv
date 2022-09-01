package cli

import (
	"github.com/gookit/gcli/v3"
	"github.com/inherelab/goenv"
	"github.com/inherelab/goenv/internal"
)

var insOpts = struct {
	dlHost string
}{}

// InstallCmd define
var InstallCmd = &gcli.Command{
	Name:    "install",
	Desc:    "install current Go version",
	Aliases: []string{"i", "ins"},
	Config: func(c *gcli.Command) {
		c.StrOpt(&insOpts.dlHost, "dl-host", "dl", "", "custom download host, default use config")

		c.AddArg("version", "want installed go version", true)
	},
	Func: func(c *gcli.Command, args []string) error {
		ver := c.Arg("version").String()

		adaptor, err := internal.NewEnvManager(goenv.Cfg.Mode).CreateAdaptor()
		if err != nil {
			return err
		}

		return adaptor.Install(ver)
	},
}
