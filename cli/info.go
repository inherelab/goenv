package cli

import (
	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/sysutil"
	"github.com/inherelab/goenv"
)

var infoOpts = struct {
	All bool
}{}

// InfoCmd define
var InfoCmd = &gcli.Command{
	Name: "info",
	Desc: "display some information",
	// Aliases: []string{"use"},
	Config: func(c *gcli.Command) {
		c.BoolOpt(&infoOpts.All, "all", "", false, "display all loaded config data")
	},
	Func: func(c *gcli.Command, args []string) error {
		color.Infoln("Go Info:")
		info, err := sysutil.OsGoInfo()
		if err != nil {
			return err
		}
		dump.NoLoc(info)

		color.Infoln("App Config:")
		dump.NoLoc(goenv.Cfg)

		if infoOpts.All {
			color.Infoln("All Config:")
			dump.NoLoc(goenv.CfgMgr.Data())
		}
		return nil
	},
}
