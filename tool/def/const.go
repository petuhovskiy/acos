package def

import "github.com/bclicn/color"

const (
	RootFile      = ".acos"
	TestsDir      = "tests"
	DefaultSource = "main.c"
	DefaultExe    = "main"

	DefaultIn  = "input.txt"
	DefaultOut = "output.txt"
	AnsSuff    = ".a"
)

var (
	OK = color.BGreen("OK")
	WA = color.BRed("WA")
)
