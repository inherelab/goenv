package internal

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v3/interact"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

// BrewAdaptor struct
type BrewAdaptor struct {
	opts *CallOpts
}

func NewBrewAdaptor() *BrewAdaptor {
	return &BrewAdaptor{
		opts: newDefaultBrewOpts(),
	}
}

func newDefaultBrewOpts() *CallOpts {
	return &CallOpts{
		LibDir: "/usr/local/opt",
	}
}

// WithOptions data
func (a *BrewAdaptor) WithOptions(opts *CallOpts) Adaptor {
	a.opts = opts
	return a
}

// List installed version
func (a *BrewAdaptor) List() error {
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
		if strings.HasPrefix(info.Version, ver) {
			indent = color.Info.Sprint("* ")
		}
		return indent + ver
	})

	color.Infoln("Installed Versions:")
	fmt.Println(strings.Join(versions, "\n"))
	return nil
}

func (a *BrewAdaptor) fmtVerAndLibPath(ver string) (string, string) {
	// 1.16.5 -> 1.16
	ver = formatVersion(ver)

	return ver, "/usr/local/opt/go@" + ver
}

// Switch go to given version
func (a *BrewAdaptor) Switch(ver string) error {
	ver, insPath := a.fmtVerAndLibPath(ver)
	if !fsutil.PathExists(insPath) {
		return errorx.Rawf("not found Go %s on %s", ver, "/usr/local/opt")
	}

	info, err := sysutil.OsGoInfo()
	if err != nil {
		return err
	}

	old := formatVersion(info.Version)
	if old == ver {
		return errorx.Rawf("The current Go version is already %s", ver)
	}

	cliutil.Infoln("Current Go version is", info.Version)
	if interact.Unconfirmed("Ensure switch to "+ver, a.opts.Yes) {
		cliutil.Infoln("Bye, Quit")
		return nil
	}

	var line string

	cmdline := "brew unlink go@" + old
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

	cliutil.Infoln("Switch successful!")
	return sysutil.NewCmd("go", "version").OutputToStd().Run()
}

// Install go by given version
func (a *BrewAdaptor) Install(ver string) error {
	ver, insPath := a.fmtVerAndLibPath(ver)
	if fsutil.PathExists(insPath) {
		return errorx.Rawf("go %s has been installed on %s", ver, insPath)
	}

	cliutil.Magentaln("Installing go", ver)

	c := sysutil.NewCmd("brew", "install")
	c.WithArgf("go@%s", ver)
	c.OutputToStd()
	c.BeforeExec = func(c *sysutil.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.Run()
}

// Update go by given version
func (a *BrewAdaptor) Update(ver string) error {
	ver, insPath := a.fmtVerAndLibPath(ver)
	if !fsutil.PathExists(insPath) {
		cliutil.Infoln("TIP: the go", ver, "not found, will be install")
		return a.Install(ver)
	}

	cliutil.Magentaln("Updating go", ver)

	c := sysutil.NewCmd("brew", "upgrade")
	c.WithArgf("go@%s", ver)
	c.OutputToStd()
	c.BeforeExec = func(c *sysutil.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.Run()
}

// Uninstall go by given version
func (a *BrewAdaptor) Uninstall(ver string) error {
	ver, insPath := a.fmtVerAndLibPath(ver)
	if !fsutil.PathExists(insPath) {
		return errorx.Rawf("not found Go %s on %s", ver, insPath)
	}

	cliutil.Magentaln("Uninstalling go", ver)

	c := sysutil.NewCmd("brew", "uninstall")
	c.WithArgf("go@%s", ver)
	c.OutputToStd()
	c.BeforeExec = func(c *sysutil.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.Run()
}

func formatVersion(ver string) string {
	// 1.16.5 -> 1.16
	ss := strings.Split(ver, ".")
	if len(ss) > 2 {
		ver = strings.Join(ss[:2], ".")
	}

	return ver
}
