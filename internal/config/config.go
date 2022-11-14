package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	bindPort string `toml:"bind_port"`
}

func NewConfig() *Config {
	conf := Config{}
	if _, err := toml.DecodeFile("configs/config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", conf.bindPort)
	return &conf
}
