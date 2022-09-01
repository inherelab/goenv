package internal

import (
	"fmt"

	"github.com/gookit/gcli/v3/interact"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/sysutil"
)

// BrewAdaptor struct
type BrewAdaptor struct {
	opts *CallOpts
}

// WithOptions data
func (a *BrewAdaptor) WithOptions(opts *CallOpts) Adaptor {
	a.opts = opts
	return a
}

func (a *BrewAdaptor) goInstallPath(ver string) string {
	return "/usr/local/opt/go@" + ver
}

// Switch go to given version
func (a *BrewAdaptor) Switch(ver string) error {
	insPath := a.goInstallPath(ver)
	if !fsutil.PathExists(insPath) {
		return errorx.Rawf("not found Go %s on %s", ver, insPath)
	}

	info, err := sysutil.OsGoInfo()
	if err != nil {
		return err
	}

	old := info.Version
	cliutil.Infoln("Current Go version is", old)
	if interact.Unconfirmed("Ensure switch to "+ver, a.opts.Yes) {
		cliutil.Infoln("Bye, Quit")
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

// Install go by given version
func (a *BrewAdaptor) Install(ver string) error {
	insPath := a.goInstallPath(ver)
	if fsutil.PathExists(insPath) {
		return errorx.Rawf("go %s has been installed on %s", ver, insPath)
	}

	cliutil.Magentaln("Installing go ", ver)

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
	insPath := a.goInstallPath(ver)
	if !fsutil.PathExists(insPath) {
		cliutil.Infoln("the go", ver, "not found, will be install")
		return a.Install(ver)
	}

	cliutil.Magentaln("Updating go ", ver)

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
	insPath := a.goInstallPath(ver)
	if !fsutil.PathExists(insPath) {
		return errorx.Rawf("not found Go %s on %s", ver, insPath)
	}

	cliutil.Magentaln("Uninstalling go ", ver)

	c := sysutil.NewCmd("brew", "uninstall")
	c.WithArgf("go@%s", ver)
	c.OutputToStd()
	c.BeforeExec = func(c *sysutil.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.Run()
}
