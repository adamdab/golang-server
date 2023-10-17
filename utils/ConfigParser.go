package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port            string `json:"port"`
	ApplicationName string `json:"application-name"`
}

func (config *Config) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(config); err != nil {
		return err
	}
	return nil
}
