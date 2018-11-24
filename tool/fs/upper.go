package fs

import (
	"fmt"
	"path/filepath"
)

func FindUpperChild(dir, child string, max int) (string, error) {
	for i := 0; i < max; i++ {
		file := Join(dir, child)
		if FileExists(file) {
			return dir, nil
		}
		dir = filepath.Dir(dir)
	}
	return "", fmt.Errorf("upper child %s not found, %d iterations done", child, max)
}
