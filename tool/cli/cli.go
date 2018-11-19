package cli

import (
	"fmt"
	"os"
)

func Start() {
	args := os.Args[1:]
	if len(args) >= 1 {
		switch args[0] {
		case "new":
			NewTask(args[1:])
		case "test":
			TestTask(args[1:])
		case "gen":
			GenTests(args[1:])
		default:
			fmt.Printf("Unknown action: %s\n", args[0])
		}
		return
	}
}
