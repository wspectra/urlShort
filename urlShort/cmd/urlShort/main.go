package main

import (
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/urlShort/internal/server"
)

func main() {
	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	//mem := store.NewInMemory()

	//fmt.Println("", mem.PostInfo("hello"))
	//fmt.Println(mem.PostInfo("world"))
	//fmt.Println(mem.GetInfo(mem.PostInfo("world")))
	//fmt.Println(mem.GetInfo(mem.PostInfo("hello")))
	//
	//fmt.Println(mem.PostInfo("world"))
	//fmt.Println(mem.GetInfo("Short"))

	//Shout(mem)
}
