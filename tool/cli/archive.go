package cli

import (
	"os"

	"github.com/petuhovskiy/acos/tool/cc"
	"github.com/petuhovskiy/acos/tool/def"
	"github.com/petuhovskiy/acos/tool/fs"

	"github.com/urfave/cli"
)

func archiveAction(c *cli.Context) error {
	g, err := def.LoadGlobal()
	if err != nil {
		cc.Errorfln("Root/config not found")
		return err
	}
	err = os.Chdir(g.Root)
	if err != nil {
		cc.Errorfln("Failed to change wd to root")
		return err
	}
	for _, task := range c.Args() {
		err := os.Rename(fs.Join(g.Tasks, task), fs.Join(g.Archive, task))
		if err != nil {
			cc.Errorfln("Failed to archive task %s", task)
			return err
		}
	}
	return nil
}

func unarchiveAction(c *cli.Context) error {
	g, err := def.LoadGlobal()
	if err != nil {
		cc.Errorfln("Root/config not found")
		return err
	}
	err = os.Chdir(g.Root)
	if err != nil {
		cc.Errorfln("Failed to change wd to root")
		return err
	}
	for _, task := range c.Args() {
		err := os.Rename(fs.Join(g.Archive, task), fs.Join(g.Tasks, task))
		if err != nil {
			cc.Errorfln("Failed to unarchive task %s", task)
			return err
		}
	}
	return nil
}
