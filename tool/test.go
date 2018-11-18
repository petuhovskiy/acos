package tool

import (
	"fmt"
	"strings"

	"github.com/bclicn/color"
)

type TestOptions struct {
	Source     string
	Executable string
	TestsDir   string
}

func TestTask(opts TestOptions) {
	Action("Compile")
	err := Compile(CompileOptions{
		Src: opts.Source,
		Dst: opts.Executable,
		Log: true,
	})
	FError("Compilation error", err)

	Action("Lint")
	Lint(opts.Source)

	Action("Test")
	tests, err := FindTests(opts.TestsDir)
	FError("Failed to find tests", err)

	for _, test := range tests {
		fmt.Println()
		fmt.Println(color.Bold(fmt.Sprintf("> Test %s", test.Name)))

		if !strings.HasSuffix(test.In, "\n") {
			FWarn("Missing \\n at " + test.InFile)
		}
		if !strings.HasSuffix(test.Ans, "\n") {
			FWarn("Missing \\n at " + test.AnsFile)
		}

		TestRun(opts.Executable, test)
	}
}
