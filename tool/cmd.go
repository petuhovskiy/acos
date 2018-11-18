package tool

import (
	"fmt"
	"os"

	"github.com/bclicn/color"
)

func Action(action string) {
	fmt.Println(color.Bold(">"), color.BLightGreen(action))
}

func FError(errorMsg string, err error) {
	if err == nil {
		return
	}
	fmt.Println()
	fmt.Println(color.Bold(">"), color.BRed(errorMsg)+color.Bold(":"), err)
	os.Exit(1)
}

func FWarn(warn string) {
	fmt.Println(color.BYellow("WARN:"), warn)
}
