package cli

import (
	"fmt"

	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/interact"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

var switchOpts = struct {
	yes  bool
	mode string
}{}

// SwitchCmd define
var SwitchCmd = &gcli.Command{
	Name:     "switch",
	Desc:     "switch current Go version",
	Examples: `${binWithCmd} -m brew 1.16`,
	Aliases:  []string{"use"},
	Config: func(c *gcli.Command) {
		c.StrOpt(&switchOpts.mode, "mode", "m", "", "allow: goenv, brew")
		c.BoolOpt(&switchOpts.yes, "yes", "y", false, "set confirm default value")

		c.AddArg("version", "the new go version for switch", true)
	},
	Func: func(c *gcli.Command, args []string) error {
		var err error
		if switchOpts.mode == "brew" {
			err = switchByBrew(c)
		} else {
			c.Infoln("TODO un-supported")
		}

		return err
	},
}

func switchByBrew(c *gcli.Command) error {
	ver := c.Arg("version").String()

	insPath := "/usr/local/opt/go@" + ver
	if !fsutil.PathExists(insPath) {
		return errorx.Rawf("not found Go %s on %s", ver, insPath)
	}

	info, err := sysutil.OsGoInfo()
	if err != nil {
		return err
	}

	old := info.Version
	c.Infoln("Current Go version is", old)
	if interact.Unconfirmed("Ensure switch to "+ver, switchOpts.yes) {
		c.Infoln("Bye, Quit")
		return nil
	}

	var line string

	cmdline := "brew unlink go"
	cliutil.Magentaln("Unbinding links for go:", cmdline)
	line, err = cliutil.ExecLine(cmdline)
	if err != nil {
		return err
	}
	fmt.Println(line)

	cmdline = "brew link go@" + ver
	cliutil.Magentaln("Binding links for go:", cmdline)
	line, err = cliutil.ExecLine(cmdline)
	if err != nil {
		return err
	}
	fmt.Println(line)

	cliutil.Successln("Switch successful, please reload shell")
	return nil
}
