package main

import (
	"github.com/kellinreaver/hafood-backend/config"
)

func main() {
	router := config.SetRouter()
	router.Run()
}
