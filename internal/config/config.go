package config

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog"
	"log"
)

var (
	ConfPath  string
	storeFlag string
)

func init() {
	flag.StringVar(&ConfPath, "config-path", "configs/config.toml", "config path")
	flag.StringVar(&storeFlag, "store-flag", "postgres", "store flag")
}

type Config struct {
	BindPort   string `toml:"bind_port"`
	DebugLevel string `toml:"debug_level"`
	Store      string `toml:"store"`

	DatabaseUrl struct {
		Host     string `toml:"host"`
		Dbname   string `toml:"dbname"`
		Sslmode  string `toml:"sslmode"`
		Password string `toml:"password"`
		User     string `toml:"user"`
	} `toml:"database_url"`
}

func NewConfig() *Config {
	flag.Parse()
	conf := Config{}
	if _, err := toml.DecodeFile(ConfPath, &conf); err != nil {
		log.Fatal(err)
	}
	conf.setLogLevel()
	if err := conf.checkingStoreFlag(); err != nil {
		log.Fatal(errors.New("[CONFIG]: wrong store flag"))
	}
	return &conf
}

func (c *Config) checkingStoreFlag() error {
	c.Store = storeFlag
	switch c.Store {
	case "inmemory":
		return nil
	case "postgres":
		return nil
	default:
		return errors.New("[CONFIG]: wrong store flag")
	}

}
func (c *Config) setLogLevel() {
	switch c.DebugLevel {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "disabled":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		log.Fatal("[CONFIG_ERROR]: wrong debug level")
	}
}
