package persistence

import "github.com/annakertesz/br-engineer-task/src/model"

type Persistence interface {
	SaveUser(user *model.User)
	SaveApp(app model.App)
	GetUser(userID string) *model.User
	GetUsers() []*model.User
	GetApp(appId string) model.App
	UpdateUser(user model.User) error
	UpdateApp(app model.App) error
	Print()
}
