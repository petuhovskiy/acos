package tool

import (
	"os"
	"os/exec"
)

type CompileOptions struct {
	Src string
	Dst string
	Log bool
}

func Compile(opts CompileOptions) error {
	cmd := exec.Command(
		"gcc",
		"-Wall",
		"-Werror",
		"-std=gnu11",
		// "-g",
		opts.Src,
		"-o",
		opts.Dst,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
