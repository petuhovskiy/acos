package cli

import (
	"github.com/petuhovskiy/acos/tool"
)

func TestTask(args []string) {
	opts := tool.TestOptions{
		Source:     "main.c",
		Executable: "./main",
		TestsDir:   "tests",
	}

	tool.TestTask(opts)
}
