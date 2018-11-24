package fs

import (
	"os"
)

func Create(path, content string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.WriteString(content)
	if err != nil {
		return err
	}

	return out.Sync()
}

func CreateDir(path string) error {
	return os.Mkdir(path, os.ModeDir|os.ModePerm)
}
