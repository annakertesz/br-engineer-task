package controller

import (
	"github.com/annakertesz/br-engineer-task/src/model"
	"time"
)

type Controller interface {
	CreateUser(userName string, plan string) model.User
	CreateApp(userID string, appName string, openSource bool) (model.App, error)
	ChangeLimits(appID string, concBuild int, buildTime time.Duration, buildPerMonth int, teamMembers int) error
	UsePrivateLimits(appID string) error
	GetLimit(appID string) model.Limit
}
