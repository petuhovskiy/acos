package tool

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bclicn/color"
)

type FileTest struct {
	Name    string
	InFile  string
	AnsFile string
	In      string
	Ans     string
}

func FindTests(dir string) ([]FileTest, error) {
	arr := make([]FileTest, 0)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		var name, aName string
		name = f.Name()
		if strings.HasSuffix(name, ".in") {
			aName = strings.TrimSuffix(name, ".in") + ".out"
		} else {
			aName = name + ".a"
		}

		stat, err := os.Stat(dir + "/" + aName)
		if err != nil || stat.Mode().IsDir() {
			continue
		}

		test := FileTest{
			Name:    name,
			InFile:  dir + "/" + name,
			AnsFile: dir + "/" + aName,
		}

		test.In, err = ReadText(test.InFile)
		if err != nil {
			fmt.Println(color.Red("Failed to read"), test.In, color.Bold("=>"), err)
			continue
		}

		test.Ans, err = ReadText(test.AnsFile)
		if err != nil {
			fmt.Println(color.Red("Failed to read"), test.Ans, color.Bold("=>"), err)
			continue
		}

		arr = append(arr, test)
	}

	return arr, nil
}
