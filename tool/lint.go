package tool

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/petuhovskiy/acos/tool/def"
)

func Lint(src string) {
	conf := def.LoadConfig()

	if conf.Defaults.AsmSrc != "" {
		err := LintAsm(conf.Defaults.AsmSrc)
		FError("Asm linter failed", err)
		return
	}

	code, err := ReadText(src)
	FError("Failed to read src", err)

	for i, line := range strings.Split(code, "\n") {
		if strings.Contains(line, "\t") {
			FWarn(fmt.Sprintf("Line %d - tab found.", i+1))
			continue
		}

		spaces := CountIndent(line)
		if spaces%4 != 0 {
			FWarn(fmt.Sprintf("Line %d - found indentation of %d spaces, must be 4n.", i+1, spaces))
			continue
		}

		if len(line) > 80 {
			FWarn(fmt.Sprintf("Line %d - too long (%d), line length must be leq 80.", i+1, len(line)))
			continue
		}

		for _, b := range line {
			if b == 127 || b < 32 && b != 10 && b != 13 {
				FWarn(fmt.Sprintf("Line %d - found control symbol, code %d", i+1, b))
			}
		}
	}

	if !strings.HasSuffix(code, "\n") {
		FWarn("Missing \\n at the end of sourcefile.")
	}
}

func LintAsm(src string) error {
	cmd := exec.Command(
		"asm_checker",
		src,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func CountIndent(line string) int {
	cnt := 0
	for ; cnt < len(line); cnt++ {
		if line[cnt] != ' ' {
			break
		}
	}
	return cnt
}
