package store

import (
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/internal/pkg/utils"
)

type InMemory struct {
	Data map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{make(map[string]string)}
}

func (inMem *InMemory) GetInfo(shortUrl string) string {
	for key, value := range inMem.Data {
		if value == shortUrl {
			return key
		}
	}
	log.Info().Msg("[INMEMORY_STORE]: Long Url not found")
	return ""
}

func (inMem *InMemory) PostInfo(longUrl string) string {
	if _, b := inMem.Data[longUrl]; b == true {
		log.Info().Msg("[INMEMORY_STORE]: short url already exists")
		return inMem.Data[longUrl]
	}
	inMem.Data[longUrl] = utils.GenerateRandomString()
	return inMem.Data[longUrl]
}
