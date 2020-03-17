package model

import "errors"

const (
	FREE  = "free"
	DEVELOPER = "developer"
	ORGANIZATION = "organization"
)

type Plan struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Limits Limit  `json:"limits"`
}

type PlanType struct {
	Free Plan `json:"free"`
	Developer Plan `json:"developer"`
	Organization Plan `json:"organization"`
}

func (plans *PlanType) Get(name string) (Plan, error) { //TODO: default:err
	switch name {
	case FREE:
		return plans.Free, nil
	case DEVELOPER:
		return plans.Developer, nil
	case ORGANIZATION:
		return plans.Organization, nil
	default:
		return Plan{}, errors.New("invalid plan name")
	}
}
