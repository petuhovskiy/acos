package tool

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bclicn/color"
)

func GenTests(gen, sol string, l, r int) {
	cnt := r - l + 1
	if cnt <= 0 || cnt > 200 {
		FError("Cnt must be in [1, 200]", fmt.Errorf("incorrent cnt: %d", cnt))
	}

	for i := l; i <= r; i++ {
		fmt.Printf("Generating test %d\n", i)

		in := fmt.Sprintf("tests/%d", i)
		ans := fmt.Sprintf("tests/%d.a", i)

		if _, err := os.Stat(in); err == nil {
			fmt.Println("Already exists, skipping...")
			continue
		}

		if _, err := os.Stat(ans); err == nil {
			fmt.Println("Already exists, skipping...")
			continue
		}

		err := GenTest(gen, in)
		if err != nil {
			fmt.Println(color.BRed("ERR:"), "test generation.", err)
			continue
		}

		err = RunSolution(sol, in, ans)
		if err != nil {
			fmt.Println(color.BRed("ERR:"), "solution run.", err)
			continue
		}
	}
}

func GenTest(exe, in string) error {
	cmd := exec.Command(exe)

	inFile, err := os.Create(in)
	if err != nil {
		return err
	}
	defer inFile.Close()

	cmd.Stdin = os.Stdin
	cmd.Stdout = inFile
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunSolution(exe, in, out string) error {
	cmd := exec.Command(exe)

	inFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdin = inFile
	cmd.Stdout = outFile
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
