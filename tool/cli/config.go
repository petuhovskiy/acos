package cli

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/petuhovskiy/acos/tool/def"
	"github.com/urfave/cli"
)

func confAction(c *cli.Context) error {
	println("Current config:")
	conf := def.LoadConfig()
	spew.Dump(conf)

	return nil
}
