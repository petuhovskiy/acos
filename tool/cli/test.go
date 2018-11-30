package cli

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/petuhovskiy/acos/tool/cc"
	"github.com/petuhovskiy/acos/tool/def"

	"github.com/petuhovskiy/acos/tool"
	"github.com/petuhovskiy/acos/tool/tests"
	"github.com/urfave/cli"
)

func testAction(c *cli.Context) error {
	src := c.String("src")
	if !strings.HasSuffix(src, ".c") {
		return errors.New("source file must ends in '.c'")
	}
	opts := tool.TestOptions{
		Source:     src,
		Executable: "./" + strings.TrimSuffix(src, ".c"),
		TestsDir:   "tests",
	}

	return tool.TestTask(opts)
}

func addTestAction(c *cli.Context) error {
	dir, err := tests.FindDir()
	if err != nil {
		return err
	}

	in, err := ioutil.ReadFile(def.DefaultIn)
	if err != nil {
		cc.Errorfln("can't read %s", def.DefaultIn)
		return err
	}

	out, err := ioutil.ReadFile(def.DefaultOut)
	if err != nil {
		cc.Errorfln("can't read %s", cc.Var(def.DefaultOut))
		return err
	}

	name, err := tests.CreateOne(dir, in, out)
	if err != nil {
		return err
	}

	cc.Okfln("Created test %s", cc.Var(name))
	return nil
}
