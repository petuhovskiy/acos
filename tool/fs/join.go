package fs

import "path/filepath"

func Join(elem ...string) string {
	return filepath.Join(elem...)
}
