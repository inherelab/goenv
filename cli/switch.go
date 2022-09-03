package cli

import (
	"strings"

	"github.com/gookit/gcli/v3"
	"github.com/inherelab/goenv/internal"
)

var switchOpts = struct {
	yes bool
}{}

// SwitchCmd define
var SwitchCmd = &gcli.Command{
	Name:     "switch",
	Desc:     "switch Go to input version",
	Examples: `${binWithCmd} 1.16`,
	Aliases:  []string{"use"},
	Config: func(c *gcli.Command) {
		c.BoolOpt(&switchOpts.yes, "yes", "y", false, "set confirm default value")

		c.AddArg("version", "the target go version for switch", true)
	},
	Func: func(c *gcli.Command, args []string) error {
		ver := c.Arg("version").String()
		ver = strings.TrimPrefix(ver, "go")

		adaptor, err := makeAdaptor()
		if err != nil {
			return err
		}

		adaptor.ApplyOpFunc(func(opts *internal.CallOpts) {
			opts.Yes = switchOpts.yes
		})
		return adaptor.Switch(ver)
	},
}
