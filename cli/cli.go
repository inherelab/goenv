package cli

import (
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/cliutil"
	"github.com/inherelab/goenv"
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
		app.Desc = "this is my cli application"
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
}

func addCommands() {
	App.Add(
		InfoCmd,
		ShellCmd,
		InstallCmd,
		SwitchCmd,
		ListCmd,
	)
}
