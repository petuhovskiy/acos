package cli

import (
	"errors"

	"github.com/petuhovskiy/acos/tool/cc"
	"github.com/petuhovskiy/acos/tool/fs"

	"github.com/urfave/cli"
)

func newAction(c *cli.Context) error {
	g, err := globalset()
	if err != nil {
		return err
	}
	for _, task := range c.Args() {
		if fs.DirExists(task) {
			cc.Errorfln("Task %s already exists", cc.Var(task))
			return errors.New("task exists")
		}
		err := fs.CreateDirCopy(fs.Join(g.Tasks, task), g.Template)
		if err != nil {
			cc.Errorfln("Failed to create task %s", cc.Var(task))
			return err
		}
		cc.Okfln("Created task %s", cc.Var(task))
	}
	return nil
}
