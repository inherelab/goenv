package cli

import (
	"github.com/gookit/gcli/v3"
)

// ListCmd define
var ListCmd = &gcli.Command{
	Name:    "list",
	Desc:    "list installed Go version",
	Aliases: []string{"ls"},
	Config: func(c *gcli.Command) {

	},
	Func: func(c *gcli.Command, args []string) error {
		adaptor, err := makeAdaptor()
		if err != nil {
			return err
		}

		return adaptor.List()
	},
}
