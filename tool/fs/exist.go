package fs

import "os"
import "github.com/petuhovskiy/acos/tool/cc"

func FileExists(file string) bool {
	stat, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		cc.Warnfln("Unknown exists error: %e", err)
		return false
	}
	if stat.IsDir() {
		cc.Warnfln("%s not file, but dir", file)
		return false
	}
	return true
}

func DirExists(dir string) bool {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		cc.Warnfln("Unknown exists error: %e", err)
		return false
	}
	if !stat.IsDir() {
		cc.Warnfln("%s not dir, but file", dir)
		return false
	}
	return true
}
