package utils

import (
	"math/rand"
	"time"
)

const (
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	alphLen     = len(alphabet)
	shortUrlLen = 10
)

func GenerateRandomString() string {

	rand.Seed(time.Now().UnixNano())
	res := make([]rune, shortUrlLen)

	for i := 0; i < shortUrlLen; i++ {
		res[i] = []rune(alphabet)[rand.Intn(alphLen)]
	}
	return string(res)
}
