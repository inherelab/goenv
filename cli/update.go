package cli

import (
	"github.com/gookit/gcli/v3"
)

var upOpts = struct {
	dlHost string
}{}

// UpdateCmd define
var UpdateCmd = &gcli.Command{
	Name:    "update",
	Desc:    "update input Go version",
	Aliases: []string{"u", "up"},
	Config: func(c *gcli.Command) {
		c.StrOpt(&upOpts.dlHost, "dl-host", "dl", "", "custom download host, default use config")

		c.AddArg("version", "want updated go version", true)
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

		return adaptor.Update(ver)
	},
}
