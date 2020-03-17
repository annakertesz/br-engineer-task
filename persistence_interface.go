package br_engineer_task

import "github.com/annakertesz/br-engineer-task/model"

type Persist interface {
	SaveUser(user *model.User)
	SaveApp(app model.App)
	GetUser(userID string) *model.User
	GetApp(appId string) model.App
	UpdateUser(user model.User)
}
