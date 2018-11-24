package fs

import "github.com/otiai10/copy"

func CreateDirCopy(dst, src string) error {
	return copy.Copy(src, dst)
}
