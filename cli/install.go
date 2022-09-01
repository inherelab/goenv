package cli

import (
	"strings"

	"github.com/gookit/gcli/v3"
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
		ver = strings.TrimPrefix(ver, "go")

		adaptor, err := makeAdaptor()
		if err != nil {
			return err
		}

		return adaptor.Install(ver)
	},
}
