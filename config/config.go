package config

import (
	"encoding/json"
	"os"
)

// Config is a struct which represents the configuration
type Config struct {
	GRPC struct {
		Host    string `json:"host"`
		Network string `json:"network"`
	} `json:"grpc"`
	HTTP struct {
		Host string `json:"host"`
	} `json:"http"`
}

// LoadConfiguration is a func which will load the configuration from config.json file.
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}
