package br_engineer_task

import "github.com/annakertesz/br-engineer-task/model"

type Controller interface {
	CreateUser(userName string, plan string) model.User
	CreateApp(userID string, isOpenSource string)
	ChangeLimits(appID string, concBuild int, buildTime int, buildPerMonth int, teamMembers int)
	OptOutLimits(appID string)
	UsePublicLimits(appID string)
	GetLimit(appID string)
}
