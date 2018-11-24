package cc

import (
	"fmt"

	"github.com/bclicn/color"
)

func Errorfln(t string, args ...interface{}) {
	fmt.Println(color.BRed(" ERR:"), fmt.Sprintf(t, args...))
}

func Warnfln(t string, args ...interface{}) {
	fmt.Println(color.BYellow("WARN:"), fmt.Sprintf(t, args...))
}
