package store

import (
	"errors"
	"github.com/wspectra/urlShort/internal/pkg/utils"
)

type InMemory struct {
	Data map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{make(map[string]string)}
}

func (inMem *InMemory) GetInfo(shortUrl string) (string, error) {
	for key, value := range inMem.Data {
		if value == shortUrl {
			return key, nil
		}
	}
	return "", errors.New("long Url not found")
}

func (inMem *InMemory) PostInfo(longUrl string) (string, error) {
	if _, b := inMem.Data[longUrl]; b == true {
		return inMem.Data[longUrl], nil
	}
	inMem.Data[longUrl] = utils.GenerateRandomString()
	return inMem.Data[longUrl], nil
}
