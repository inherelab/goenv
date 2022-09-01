package cli

import "github.com/gookit/gcli/v3"

// ShellCmd define
var ShellCmd = &gcli.Command{
	Name: "shell",
	Desc: "generate shell script codes",
	Help: `
${binWithCmd} shell
`,
	Aliases: []string{"script"},
	Config: func(c *gcli.Command) {

	},
	Func: func(c *gcli.Command, args []string) error {
		return nil
	},
}
