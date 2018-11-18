package cli

import (
	"fmt"
	"io"
	"os"
)

func NewTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Specify task name")
		os.Exit(1)
	}

	name := args[0]
	err := os.Mkdir(name, os.ModeDir|os.ModePerm)
	if os.IsExist(err) {
		fmt.Println("Task already exists")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Mkdir(name+"/tests", os.ModeDir|os.ModePerm)

	silentCopy("z/main.c", name+"/main.c")

	fmt.Println("Created new task.")
}

func silentCopy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Sync()
}
