package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/internal/server"
	"github.com/wspectra/urlShort/internal/store"
	"strings"
)

func Shout(a store.Store) {
	fmt.Println(strings.ToUpper(a.PostInfo("biba")))

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
