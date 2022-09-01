package cli

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/sysutil"
	"github.com/inherelab/goenv"
)

// ListCmd define
var ListCmd = &gcli.Command{
	Name:    "list",
	Desc:    "list current Go version",
	Aliases: []string{"ls"},
	Config: func(c *gcli.Command) {

	},
	Func: func(c *gcli.Command, args []string) error {
		info, err := sysutil.OsGoInfo()
		if err != nil {
			return err
		}

		str, err := sysutil.ShellExec("ls /usr/local/opt | grep go@")
		if err != nil {
			return errorx.Wrap(err, "find local go error")
		}

		lines := strings.Split(strings.TrimSpace(str), "\n")

		prefix := "go@"
		versions := arrutil.StringsMap(lines, func(s string) string {
			ver := strings.TrimPrefix(s, prefix)

			indent := "  "
			if ver == info.Version {
				indent = color.Info.Sprint("* ")
			} else if goenv.Cfg.IsBrewMode() && strings.HasPrefix(info.Version, ver) {
				indent = color.Info.Sprint("* ")
			}
			return indent + ver
		})

		color.Infoln("Installed Versions:")
		fmt.Println(strings.Join(versions, "\n"))
		return nil
	},
}
