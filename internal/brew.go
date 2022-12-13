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
	"github.com/gookit/goutil/sysutil/cmdr"
)

// BrewAdaptor struct
type BrewAdaptor struct {
	baseAdaptor
}

// NewBrewAdaptor instance
func NewBrewAdaptor() *BrewAdaptor {
	return &BrewAdaptor{
		baseAdaptor: newBaseAdaptor(AdaptorBrew, newDefaultBrewOpts()),
	}
}

func newDefaultBrewOpts() *CallOpts {
	return &CallOpts{
		LibDir: "/usr/local/opt",
	}
}

// List installed version
func (a *BrewAdaptor) List() error {
	info, err := sysutil.OsGoInfo()
	if err != nil {
		return err
	}

	prefix := "go@"
	cmdline := fmt.Sprintf("ls %s | grep %s", a.opts.LibDir, prefix)
	fmt.Println("Find local Go version at", a.opts.LibDir)

	str, err := sysutil.ShellExec(cmdline)
	if err != nil {
		return errorx.Wrap(err, "find local go error")
	}

	lines := strings.Split(strings.TrimSpace(str), "\n")
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
	ver = formatBrewVersion(ver)

	return ver, "/usr/local/opt/go@" + ver
}

// Switch go to given version
func (a *BrewAdaptor) Switch(ver string) error {
	ver, insPath := a.fmtVerAndLibPath(ver)
	if !fsutil.PathExists(insPath) {
		if interact.Confirm(fmt.Sprintf("Not found go%s. Install now?", ver), true) {
			return a.Install(ver)
		}

		return errorx.Rawf("not found Go %s on %s", ver, a.opts.LibDir)
	}

	info, err := sysutil.OsGoInfo()
	if err != nil {
		return err
	}

	old := formatBrewVersion(info.Version)
	if old == ver {
		return errorx.Rawf("The current Go version is already %s", ver)
	}

	cliutil.Infoln("Current Go version is", info.Version)
	if interact.Unconfirmed("Ensure switch to "+ver, a.opts.Yes) {
		cliutil.Infoln("Bye, Quit")
		return nil
	}

	// cmdline = "brew unlink go@" + old
	cmdArgs := arrutil.Strings{"brew", "unlink", "go@" + old}
	cliutil.Magentaln("Unbinding links for go:", cmdArgs.Join(" "))
	err = sysutil.FlushExec(cmdArgs[0], cmdArgs[1:]...)
	if err != nil {
		return err
	}
	fmt.Println()

	// "brew link go@" + ver
	cmdArgs = arrutil.Strings{"brew", "link", "go@" + ver}
	cliutil.Magentaln("Binding links for go:", cmdArgs.Join(" "))
	err = sysutil.FlushExec(cmdArgs[0], cmdArgs[1:]...)
	if err != nil {
		return err
	}

	cliutil.Infoln("\nSwitch successful!")
	color.Bold.Print("Current: ")
	return sysutil.NewCmd("go", "version").FlushRun()
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
	c.BeforeRun = func(c *cmdr.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.FlushRun()
}

// Update go by given version
func (a *BrewAdaptor) Update(ver string) error {
	ver, insPath := a.fmtVerAndLibPath(ver)
	if !fsutil.PathExists(insPath) {
		cliutil.Infoln("TIP: the go", ver, "not found, will be install now")
		return a.Install(ver)
	}

	cliutil.Magentaln("Updating go", ver)

	c := sysutil.NewCmd("brew", "upgrade")
	c.WithArgf("go@%s", ver)
	c.BeforeRun = func(c *cmdr.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.FlushRun()
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
	c.BeforeRun = func(c *cmdr.Cmd) {
		cliutil.Yellowln(">", c.Cmdline())
	}

	return c.FlushRun()
}

func formatBrewVersion(ver string) string {
	// 1.16.5 -> 1.16
	ss := strings.Split(ver, ".")
	if len(ss) > 2 {
		ver = strings.Join(ss[:2], ".")
	}

	return ver
}
