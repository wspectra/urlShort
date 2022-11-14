package main

import (
	"fmt"
	"github.com/wspectra/urlShort/internal/config"
	"github.com/wspectra/urlShort/internal/server"
	"log"
)

func main() {
	conf := config.NewConfig()
	fmt.Println(conf)
	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Fatal("server error")
	}

}
