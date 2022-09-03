package cli

import (
	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/show"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
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
		c.BoolOpt(&infoOpts.All, "all", "a", false, "display all loaded config data")
		c.BoolOpt(&infoOpts.Init, "init", "ri", false, "re-init goenv config to user config dir")
	},
	Func: func(c *gcli.Command, args []string) error {
		if infoOpts.Init {
			_, err := goenv.InitConfigFile(true)
			return err
		}

		color.Cyanln(strutil.Repeat(" ", 15), "Goenv Information")
		color.Cyanln(strutil.Repeat("=", 45))
		info, err := sysutil.OsGoInfo()
		if err != nil {
			return err
		}

		show.AList("GO Info", info)
		show.AList("App Config", goenv.Cfg)

		if infoOpts.All {
			color.Infop("\nAll Config: ")
			dump.NoLoc(goenv.CfgMgr.Data())
		}
		return nil
	},
}
