package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Plan struct {
	Free Limit `json:"free"`
	Developer Limit `json:"developer"`
	Organization Limit `json:"organization"`
	OpensourceDefault Limit `json:"opensource_default"`
}

func GetPlansFromConfig(file string) Plan {
	var plan Plan
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &plan)
	return plan
}
