package tool

import (
	"os"
	"os/exec"

	"github.com/petuhovskiy/acos/tool/def"
)

type CompileOptions struct {
	Src string
	Dst string
	Log bool
}

func Compile(opts CompileOptions) error {
	conf := def.LoadConfig()

	args := append([]string{}, conf.Defaults.CompileArgs...)

	for i, v := range args {
		if v == "$src" {
			args[i] = opts.Src
		} else if v == "$dst" {
			args[i] = opts.Dst
		}
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
