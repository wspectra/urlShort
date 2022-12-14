package main

import (
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/internal/server"
)

func main() {
	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
