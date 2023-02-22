package api

import (
	"math/rand"
	"time"
)

func RecommandVtuber(reqmessage string) (message string) {

	rand.Seed(time.Now().UnixNano())
	if reqmessage == "hello" {
		return "hello"
	}
	return reqmessage
}
