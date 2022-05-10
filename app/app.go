package app

import (
	"log"
	"os"

	"github.com/gustavoteixeira8/cli-golang/app/commands"
	"github.com/urfave/cli"
)

var app *cli.App

func DefineCommands() {
	hostFlag := cli.StringFlag{
		Name:  "host",
		Value: "github.com",
	}

	app.Commands = []cli.Command{
		{
			Name:   "search-ip",
			Usage:  "Searches IP addresses on the internet",
			Flags:  []cli.Flag{hostFlag},
			Action: commands.SearchIps,
		},
		{
			Name:   "search-server",
			Usage:  "Searches the name of servers on the internet",
			Flags:  []cli.Flag{hostFlag},
			Action: commands.SearchServer,
		},
	}
}

func CreateApp() *cli.App {
	app = cli.NewApp()
	app.Name = "CLI-Application"
	app.Usage = "CLI application for searching IPs and server names"

	DefineCommands()

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error => %v", err)
	}

	return app
}
