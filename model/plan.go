package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	FREE  = "free"
	DEVELOPER = "developer"
	ORGANIZATION = "organization"
)

type Plan struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Limits Limit `json:"limits"`
}

type PlanType struct {
	Free Plan `json:"free"`
	Developer Plan `json:"developer"`
	Organization Plan `json:"organization"`
}

func (plans *PlanType) Get(name string) Plan { //TODO: default:err
	switch name {
	case FREE:
		return plans.Free
	case DEVELOPER:
		return plans.Developer
	case ORGANIZATION:
		return plans.Organization
	}
	return Plan{
	}
}

func GetPlansFromConfig(file string) PlanType {
	var plan PlanType
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &plan)
	return plan
}
