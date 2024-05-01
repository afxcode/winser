package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

var configFile = "winser_config.json"

type Config struct {
	PinedServiceNames []string `json:"pined_service_names"`
}

func readConfig() (conf Config, err error) {
	if _, e := os.Stat(configFile); errors.Is(e, os.ErrNotExist) {
		file, _ := json.MarshalIndent(conf, "", "    ")
		if err = os.WriteFile(configFile, file, 0644); err != nil {
			return
		}
	}

	jsonFile, err := os.Open(configFile)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteValue, &conf)
	return
}

func updateConfig() (err error) {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}

	conf := Config{}
	if err = json.Unmarshal(byteValue, &conf); err != nil {
		return
	}

	conf.PinedServiceNames = getPinedServiceNames()

	newJSON, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return
	}

	if err = os.WriteFile(configFile, newJSON, os.ModePerm); err != nil {
		return
	}
	return
}
