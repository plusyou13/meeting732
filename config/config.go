package config

import (
	"flag"
	"fmt"
	"github.com/plusyou13/comm-go/com_conf"
)

var (
	gCfg *Config
)

func GetConfig() *Config {
	return gCfg
}

type Config struct {
	Port int `yaml:"port"`

	Log struct {
		LogPath  string `yaml:"log_path"`
		LogSize  int64  `yaml:"log_size"`
		LogCount int    `yaml:"log_count"`
		LogLevel int    `yaml:"log_level"`
	} `yaml:"log"`

	MongoDB struct {
		Host     string `yaml:"host"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mongodb"`
}

func init() {
	var (
		confPath *string = flag.String("conf", "./settings.yml", "--config file")
	)
	flag.Parse()

	var err error
	gCfg = &Config{}
	if err = com_conf.YamlFromFile(&gCfg, *confPath); err != nil {
		panic(err)
	}
	fmt.Printf("cfg:%+v\n", gCfg)
}

