package config

import (
	"encoding/json"
	"os"
)

type Config struct { 
	Chance string `json:"chance"` //read
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err !=nil {
		return config, err
	}
	jsonParser :=json.NewDecoder(configFile)
	err =jsonParser.Decode(&config)
	return config, err
}