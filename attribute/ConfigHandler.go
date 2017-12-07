package attribute

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	database struct {
		connectionString string
	}
}

func GetEnvConfig() Configuration {
	file, _ := os.Open("env.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
