package persistence

import (
	"github.com/annakertesz/br-engineer-task/src/model"
	"github.com/lithammer/shortuuid"
)

type DumbPersistence struct {
	Users []*model.User
	Apps  []model.App
}

func NewDumbPersistence() *DumbPersistence {
	return &DumbPersistence{
		Users: make([]*model.User, 0),
	}
}

func (p *DumbPersistence) SaveUser(user *model.User) {
	user.SetId(shortuuid.New())
	p.Users = append(p.Users, user)
}

func (p *DumbPersistence) SaveApp(app model.App) {
	app.SetId(shortuuid.New())
	p.Apps = append(p.Apps, app)
	p.UpdateUser(*app.GetUser())
}

func (p *DumbPersistence) GetUsers() []*model.User {
	return p.Users
}

func (p *DumbPersistence) GetApp(appId string) model.App {
	for _, app := range p.Apps {
		if app.GetId() == appId {
			return app
		}
	}
	return nil
}

func (p *DumbPersistence) GetUser(userId string) *model.User {
	for _, user := range p.Users {
		if user.GetId() == userId {
			return user
		}
	}
	return nil
}

func (p *DumbPersistence) UpdateUser(user model.User) { //TODO: handle if ID doesnt exists
	persistedUser := p.GetUser(user.GetId())
	*persistedUser = user
}

func (p *DumbPersistence) UpdateApp(app model.App) { //TODO: handle if ID doesnt exists
	for i := range p.Apps {
		if p.Apps[i].GetId() == app.GetId() {
			p.Apps[i]=app
		}
	}
}