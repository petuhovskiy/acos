package cli

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func Start() {
	app := cli.NewApp()
	app.Name = "acos"
	app.Usage = "Tool for testing acos tasks"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "init in current directory",
			Action: initAction,
		},
		{
			Name:    "archive",
			Aliases: []string{"a"},
			Usage:   "archive task: tasks->archive",
			Action:  archiveAction,
		},
		{
			Name:    "unarchive",
			Aliases: []string{"u"},
			Usage:   "unarchive task: tasks<-archive",
			Action:  unarchiveAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
