package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog"
	"log"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "config-path", "configs/config.toml", "config path")
}

type Config struct {
	BindPort   string `toml:"bind_port"`
	DebugLevel string `toml:"debug_level"`
	Store      string `toml:"store"`
}

func NewConfig() *Config {
	flag.Parse()
	conf := Config{}
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		log.Fatal(err)
	}
	conf.setLogLevel()
	return &conf
}

func (c *Config) setLogLevel() {
	switch c.DebugLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		log.Fatal("[CONFIG_ERROR]: wrong debug level")
	}
}
