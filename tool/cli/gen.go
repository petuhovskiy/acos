package cli

import (
	"log"
	"strconv"

	"github.com/petuhovskiy/acos/tool"
)

func GenTests(args []string) {
	if len(args) != 4 {
		log.Fatal("== 4 arguments")
	}

	exe := args[0]
	sol := args[1]

	l, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	r, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatal(err)
	}

	tool.GenTests(exe, sol, l, r)
}
