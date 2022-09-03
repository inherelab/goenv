package internal

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/sysutil"
	"github.com/inherelab/goenv"
)

// ScoopAdaptor struct
type ScoopAdaptor struct {
	opts *CallOpts
}

// NewScoopAdaptor instance
func NewScoopAdaptor() *ScoopAdaptor {
	return &ScoopAdaptor{
		opts: newDefaultScoopOpts(),
	}
}

func newDefaultScoopOpts() *CallOpts {
	return &CallOpts{
		LibDir: "/usr/local/go",
	}
}

// Name of adaptor
func (a *ScoopAdaptor) Name() string {
	return goenv.ModeScoop
}

// ApplyOpFunc handle
func (a *ScoopAdaptor) ApplyOpFunc(fn OpFunc) {
	fn(a.opts)
}

func (a *ScoopAdaptor) List() error {
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

func (a *ScoopAdaptor) Switch(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *ScoopAdaptor) Install(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *ScoopAdaptor) Update(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *ScoopAdaptor) Uninstall(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}
