package def

import (
	"encoding/json"
	"log"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-config/source/memory"
)

type Config struct {
	Defaults Defaults
}

type Defaults struct {
	Source      string
	Exe         string
	CompileArgs []string
	RunArgs     []string
	AsmSrc      string
}

var defaultConfig = Config{
	Defaults: Defaults{
		Source: "main.c",
		Exe:    "main",
		CompileArgs: []string{
			"gcc",
			"-Wall",
			"-Werror",
			"-std=gnu11",
			"-lm",
			"-O2",
			// "-g",
			"$src",
			"-o",
			"$dst",
		},
		RunArgs: []string{
			"$exe",
		},
		AsmSrc: "",
	},
}

func LoadConfig() *Config {
	conf := config.NewConfig()

	data, err := json.Marshal(defaultConfig)
	if err != nil {
		log.Println(err)
	}

	err = conf.Load(
		memory.NewSource(
			memory.WithChangeSet(
				&source.ChangeSet{
					Data: data,
				},
			),
		),
		file.NewSource(
			file.WithPath("acos.toml"),
		),
	)
	if err != nil {
		log.Println(err)
	}

	c := &Config{}
	err = conf.Scan(&c)
	if err != nil {
		log.Println(err)
	}

	return c
}
