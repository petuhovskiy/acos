package tests

import (
	"errors"
	"fmt"

	"github.com/petuhovskiy/acos/tool/def"
	"github.com/petuhovskiy/acos/tool/fs"
)

func TestExists(dir, name string) bool {
	return fs.FileExists(fs.Join(dir, name)) ||
		fs.FileExists(fs.Join(dir, name+def.AnsSuff))
}

func FindNewTestName(dir string) (string, error) {
	l := 1
	r := 1 << 16
	res := ""
	for l <= r {
		mid := (l + r) / 2
		name := fmt.Sprintf("%02d", mid)
		if TestExists(dir, name) {
			l = mid + 1
		} else {
			res = name
			r = mid - 1
		}
	}
	if res != "" {
		return res, nil
	}
	return "", errors.New("too many tests")
}

func CreateOne(dir string, in, out []byte) (string, error) {
	name, err := FindNewTestName(dir)
	if err != nil {
		return "", err
	}

	err = fs.WriteFile(fs.Join(dir, name), in)
	if err != nil {
		return "", err
	}

	err = fs.WriteFile(fs.Join(dir, name+def.AnsSuff), out)
	if err != nil {
		return "", err
	}

	return name, nil
}
