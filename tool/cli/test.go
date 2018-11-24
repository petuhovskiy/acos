package cli

import (
	"github.com/petuhovskiy/acos/tool"
	"github.com/urfave/cli"
)

func testAction(c *cli.Context) error {
	opts := tool.TestOptions{
		Source:     "main.c",
		Executable: "./main",
		TestsDir:   "tests",
	}

	return tool.TestTask(opts)
}
