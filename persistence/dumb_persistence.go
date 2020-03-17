package persistence

import (
	"github.com/annakertesz/br-engineer-task/model"
	"github.com/lithammer/shortuuid"
)

type DumbPersistence struct {
	users []*model.User
	apps []*model.App
}

func NewDumbPersistence() *DumbPersistence {
	return &DumbPersistence{
		users: make([]*model.User, 0),
	}
}

func (p *DumbPersistence) SaveUser(user *model.User) {
	user.SetId(shortuuid.New())
	p.users = append(p.users, user)
}

func (p *DumbPersistence) SaveApp(app model.App) {
	app.SetId(shortuuid.New())
	p.apps = append(p.apps, &app)
	p.UpdateUser(*app.GetUser())
}

func (p *DumbPersistence) GetUsers() []*model.User {
	return p.users
}

func (p *DumbPersistence) GetUser(userId string) *model.User {
	for _, user := range p.users {
		if user.GetId() == userId {
			return user
		}
	}
	return nil
}

func (p *DumbPersistence) UpdateUser(user model.User) {  //TODO: handle if ID doesnt exists
	persistedUser := p.GetUser(user.GetId())
	*persistedUser = user
}
