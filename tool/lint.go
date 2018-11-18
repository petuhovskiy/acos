package tool

import (
	"fmt"
	"strings"
)

func Lint(src string) {
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
	}

	if !strings.HasSuffix(code, "\n") {
		FWarn("Missing \\n at the end of sourcefile.")
	}
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
