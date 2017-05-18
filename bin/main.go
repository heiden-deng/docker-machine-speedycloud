package main

import (
	"os"
	"path"
	"github.com/urfave/cli"
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/wujie4/docker-machine-speedycloud"
)

var appHelpTemplate = `This is a Docker Machine plugin for SpeedyCloud.
Plugin binaries are not intended to be invoked directly.
Please use this plugin through the main 'docker-machine' binary.

Version: {{.Version}}{{if or .Author .Email}}

Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
`

func main() {
	cli.AppHelpTemplate = appHelpTemplate
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "This is a Docker Machine plugin binary. Please use it through the main 'docker-machine' binary."
	app.Author = "wujie"
	app.Email = "wujieyhy2016@gmail.com"
	app.Version = speedycloud.FullVersion()
	app.Action = func(c *cli.Context) {
		plugin.RegisterDriver(speedycloud.NewDriver("", ""))
	}

	app.Run(os.Args)
}
