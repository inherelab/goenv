package cli

import (
	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/sysutil"
	"github.com/inherelab/goenv"
)

var infoOpts = struct {
	All  bool
	Init bool // init goenv config
}{}

// InfoCmd define
var InfoCmd = &gcli.Command{
	Name:    "info",
	Desc:    "display some information",
	Aliases: []string{"show"},
	Config: func(c *gcli.Command) {
		c.BoolOpt(&infoOpts.All, "all", "", false, "display all loaded config data")
		c.BoolOpt(&infoOpts.Init, "init", "", false, "re-init goenv config to user config dir")
	},
	Func: func(c *gcli.Command, args []string) error {
		if infoOpts.Init {
			_, err := goenv.InitConfigFile(true)
			return err
		}

		color.Magentaln("Goenv Information")

		color.Infop("\nGo Info: ")
		info, err := sysutil.OsGoInfo()
		if err != nil {
			return err
		}
		dump.NoLoc(info)

		color.Infop("\nApp Config: ")
		dump.NoLoc(goenv.Cfg)

		if infoOpts.All {
			color.Infop("\nAll Config: ")
			dump.NoLoc(goenv.CfgMgr.Data())
		}
		return nil
	},
}
