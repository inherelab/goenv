package cli

import (
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/cliutil"
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

func createApp() {
	App = gcli.NewApp(func(app *gcli.App) {
		app.Version = goenv.Version
		app.Desc = "Go multi version env manager"
		app.On(gcli.EvtAppInit, func(data ...interface{}) bool {
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
	gcli.GOpts().SetDisable()

	// App.GOptsBinder = func(gf *gcli.Flags) {
	// 	gf.BoolOpt()
	// }
}

func addCommands() {
	App.Add(
		InfoCmd,
		ShellCmd,
		UpdateCmd,
		InstallCmd,
		UninstallCmd,
		SwitchCmd,
		ListCmd,
	)
}

// MakeAdaptor instance
func makeAdaptor() (internal.Adaptor, error) {
	return internal.NewEnvManager(goenv.Cfg.Mode).MakeAdaptor()
}
