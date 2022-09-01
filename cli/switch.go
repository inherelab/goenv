package cli

import (
	"github.com/gookit/gcli/v3"
	"github.com/inherelab/goenv"
	"github.com/inherelab/goenv/internal"
)

var switchOpts = struct {
	yes bool
}{}

// SwitchCmd define
var SwitchCmd = &gcli.Command{
	Name:     "switch",
	Desc:     "switch current Go to given version",
	Examples: `${binWithCmd} 1.16`,
	Aliases:  []string{"use"},
	Config: func(c *gcli.Command) {
		c.BoolOpt(&switchOpts.yes, "yes", "y", false, "set confirm default value")

		c.AddArg("version", "the target go version for switch", true)
	},
	Func: func(c *gcli.Command, args []string) error {
		ver := c.Arg("version").String()

		adaptor, err := internal.NewEnvManager(goenv.Cfg.Mode).CreateAdaptor()
		if err != nil {
			return err
		}

		opts := &internal.CallOpts{
			Yes: switchOpts.yes,
		}

		return adaptor.WithOptions(opts).Switch(ver)
	},
}
