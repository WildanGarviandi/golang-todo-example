package main

import (
	"github.com/WildanGarviandi/hafood-backend/config"
)

func main() {
	router := config.SetRouter()
	router.Run(":8080")
}
