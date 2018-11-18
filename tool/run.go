package tool

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bclicn/color"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestRun(exe string, test FileTest) {
	cmd := exec.Command("./" + exe)
	out := bytes.Buffer{}
	cmd.Stdin = strings.NewReader(test.In)
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(color.BRed("ERR:"), "Solution failed.", err)
		return
	}

	outString := string(out.Bytes())
	if !strings.HasSuffix(outString, "\n") {
		FWarn("Missing \\n")
	}

	if outString == test.Ans {
		fmt.Println(color.BGreen("OK"))
		return
	}

	fmt.Println(color.BRed("WA"), "Ans != Out")

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(outString, test.Ans, false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
