package main

import (
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/internal/server"
	"github.com/wspectra/urlShort/internal/store"
)

func Shout(a store.Store) {

}

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
