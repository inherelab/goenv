package cli

import (
	"strings"

	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/events"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/strutil"
	"github.com/inherelab/goenv"
	"github.com/inherelab/goenv/internal"
)

// App instance
var App *gcli.App

// Run cli app
func Run() {
	err := goenv.Init()
	if err != nil {
		cliutil.Errorln("ERROR", err)
		return
	}

	createApp()

	addCommands()

	App.Run(nil)
}

// re-init config to file
var reInit bool

func createApp() {
	App = gcli.NewApp(func(app *gcli.App) {
		app.Version = goenv.Version
		app.Desc = "Go multi version env manager"
		app.On(gcli.EvtAppInit, func(ctx *gcli.HookCtx) bool {
			// do something...
			// fmt.Println("init app")
			return false
		})

		// app.SetVerbose(gcli.VerbDebug)
		// app.DefaultCommand("example")
		app.Logo.Text = `   ________    _______
  / ____/ /   /  _/   |  ____  ____
 / /   / /    / // /| | / __ \/ __ \
/ /___/ /____/ // ___ |/ /_/ / /_/ /
\____/_____/___/_/  |_/ .___/ .___/
                     /_/   /_/`
	})

	// disable global options
	App.Opts().SetDisable()

	App.On(events.OnAppBindOptsAfter, func(ctx *gcli.HookCtx) (stop bool) {
		ctx.App.Flags().
			BoolOpt(&reInit, "reinit", "ri", false, "re-init goenv config to user config dir")
		return false
	})

	App.On(events.OnAppOptsParsed, func(ctx *gcli.HookCtx) bool {
		if reInit {
			_, err := goenv.InitConfigFile(true)
			if err != nil {
				ctx.SetStop(true)
			}
		}

		return ctx.Stopped()
	})
}

func addCommands() {
	App.Add(
		InfoCmd,
		// ShellCmd,
		UpdateCmd,
		InstallCmd,
		UninstallCmd,
		SwitchCmd,
		ListCmd,
	)
}

// MakeAdaptor instance
func makeAdaptor() (internal.Adaptor, error) {
	return internal.MakeAdaptor(goenv.Cfg.Adaptor)
}

func checkFormatVersion(ver string) (string, error) {
	ver = strings.TrimPrefix(ver, "go")

	if !strutil.IsVersion(ver) {
		return "", errorx.Rawf("invalid version string: %s", ver)
	}
	return ver, nil
}
