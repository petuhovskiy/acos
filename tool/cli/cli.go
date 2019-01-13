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
			Usage:  "Init in current directory",
			Action: initAction,
		},
		{
			Name:    "archive",
			Aliases: []string{"a"},
			Usage:   "Archive task",
			Action:  archiveAction,
		},
		{
			Name:    "unarchive",
			Aliases: []string{"u"},
			Usage:   "Unarchive task",
			Action:  unarchiveAction,
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create new task",
			Action:  newAction,
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Test source code",
			Action:  testAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "src, s",
					Value: "main.c",
					Usage: "Compile source `FILE`",
				},
				cli.StringFlag{
					Name:  "tests, t",
					Value: "tests",
					Usage: "Tests `DIR`",
				},
			},
		},
		{
			Name:    "addtest",
			Aliases: []string{"at"},
			Usage:   "Add test from input/output.txt",
			Action:  addTestAction,
		},
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "WIP",
			Action:  genAction,
		},
		{
			Name:   "conf",
			Usage:  "Print current loaded config",
			Action: confAction,
		},
		{
			Name: "compile",
			Usage: "Just compile",
			Action: compileAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
