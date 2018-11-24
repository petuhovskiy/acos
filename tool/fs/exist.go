package fs

import "os"
import "github.com/petuhovskiy/acos/tool/cc"

func FileExists(file string) bool {
	stat, err := os.Stat(file)
	if err != nil {
		return false
	}
	if stat.IsDir() {
		cc.Warnfln("%s not file, but dir", file)
	}
	return true
}
