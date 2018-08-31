package config

import (
	"log"
	"os"
	"encoding/json"
)

type Configuration struct{
	Port		string
	Mongo_Url	string
}

func Config() Configuration{
	env := os.Getenv("env")
	configPath := "config/config." + env + ".json"
	file, _ := os.Open(configPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}

	return configuration
}