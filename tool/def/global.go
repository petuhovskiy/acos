package def

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/petuhovskiy/acos/tool/cc"
	"github.com/petuhovskiy/acos/tool/fs"
)

type Global struct {
	Config string
	Root   string
	RootConfig
}

type RootConfig struct {
	Tasks    string
	Archive  string
	Template string
}

func LoadGlobal() (Global, error) {
	dir, err := os.Getwd()
	if err != nil {
		return Global{}, err
	}

	dir, err = fs.FindUpperChild(dir, RootFile, 5)
	if err != nil {
		cc.Errorfln("acos root directory not found")
		return Global{}, err
	}

	g := Global{
		Root:   dir,
		Config: fs.Join(dir, RootFile),
	}

	configContent, err := ioutil.ReadFile(g.Config)
	if err != nil {
		return g, err
	}

	_, err = toml.Decode(string(configContent), &g.RootConfig)
	if err != nil {
		return g, err
	}

	return g, nil
}
