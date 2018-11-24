package cli

import (
	"github.com/petuhovskiy/acos/tool/cc"
	"github.com/petuhovskiy/acos/tool/def"
	"github.com/petuhovskiy/acos/tool/fs"
	"github.com/urfave/cli"
)

const (
	templateDir = "z"
	archiveDir  = "zarch"
	rootContent = `
Template = "` + templateDir + `"
Archive = "` + archiveDir + `"
Root = "."
Tasks = "."
`
	cppContent = `#include <cstdio.h>

int main()
{
}
`
)

func initAction(c *cli.Context) error {
	err := fs.Create(def.RootFile, rootContent)
	if err != nil {
		cc.Errorfln("Failed to init root config %s", def.RootFile)
		return err
	}

	err = initTemplate(templateDir)
	if err != nil {
		cc.Errorfln("Failed to init template dir in %s", templateDir)
		return err
	}

	err = initArchive(archiveDir)
	if err != nil {
		cc.Errorfln("Failed to init archive dir in %s", archiveDir)
		return err
	}
	return nil
}

func initTemplate(dir string) error {
	err := fs.CreateDir(dir)
	if err != nil {
		return err
	}

	err = fs.CreateDir(fs.Join(dir, def.TestsDir))
	if err != nil {
		return err
	}

	err = fs.Create(fs.Join(dir, def.DefaultSource), cppContent)
	if err != nil {
		return err
	}

	return nil
}

func initArchive(dir string) error {
	return fs.CreateDir(dir)
}
