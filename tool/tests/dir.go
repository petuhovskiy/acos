package tests

import (
	"errors"

	"github.com/petuhovskiy/acos/tool/fs"
)

func FindDir() (string, error) {
	dir := "./tests"
	if !fs.DirExists(dir) {
		return "", errors.New("tests dir not found")
	}
	return dir, nil
}
