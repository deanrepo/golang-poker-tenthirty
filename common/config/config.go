package config

import (
	"encoding/json"
	"os"
)

// Config represents configuration files.
type Config struct {
	ListenAddress string `json:"listenAddress"`
	RelationalDB  string `josn:"relationalDB"`
	RelationalDSN string `json:"relationalDSN"`
}

// LoadConfig loads configurations from file.
func LoadConfig(filename string) (config *Config, err error) {
	configFile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer configFile.Close()

	config = &Config{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(config)
	return
}
