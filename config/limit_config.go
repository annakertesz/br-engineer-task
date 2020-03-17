package config

import (
	"encoding/json"
	"github.com/annakertesz/br-engineer-task/model"
	"io/ioutil"
	"os"
)

type Config struct {
	OpensourceDefault model.Limit `json:"opensource_default"`
	Plans model.PlanType `json:"plans"`
}

func GetConfigFromFile(file string) (*Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, &config)
	if err!= nil {
		return nil, err
	}
	return &config, nil
}

