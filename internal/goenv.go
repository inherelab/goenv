package internal

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/sysutil"
)

// GoEnvAdaptor struct
type GoEnvAdaptor struct {
	baseAdaptor
}

// NewGoEnvAdaptor instance
func NewGoEnvAdaptor() *GoEnvAdaptor {
	return &GoEnvAdaptor{
		baseAdaptor: newBaseAdaptor(AdaptorGoEnv, newDefaultGoEnvOpts()),
	}
}

func newDefaultGoEnvOpts() *CallOpts {
	return &CallOpts{
		LibDir: "/usr/local/go",
	}
}

func (a *GoEnvAdaptor) List() error {
	info, err := sysutil.OsGoInfo()
	if err != nil {
		return err
	}

	cmdline := "ls " + a.opts.LibDir
	str, err := sysutil.ShellExec(cmdline)
	if err != nil {
		return errorx.Wrap(err, "find local go error")
	}

	lines := strings.Split(strings.TrimSpace(str), "\n")

	versions := arrutil.StringsMap(lines, func(ver string) string {
		indent := "  "
		if ver == info.Version {
			indent = color.Info.Sprint("* ")
		}
		return indent + ver
	})

	color.Infoln("Installed Versions:")
	fmt.Println(strings.Join(versions, "\n"))
	return nil
}

func (a *GoEnvAdaptor) Switch(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *GoEnvAdaptor) Install(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *GoEnvAdaptor) Update(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *GoEnvAdaptor) Uninstall(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}
