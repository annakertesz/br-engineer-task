package persistence

import (
	"errors"
	"fmt"
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
		Apps: make([]model.App, 0),
	}
}

func (p *DumbPersistence) SaveUser(user *model.User){
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

func (p *DumbPersistence) UpdateUser(user model.User) error { //TODO: handle if ID doesnt exists
	persistedUser := p.GetUser(user.GetId())
	if persistedUser == nil {
		return errors.New("Wrong userID")
	}
	*persistedUser = user
	return nil
}

func (p *DumbPersistence) UpdateApp(app model.App) error { //TODO: handle if ID doesnt exists
	for i := range p.Apps {
		if p.Apps[i].GetId() == app.GetId() {
			p.Apps[i]=app
			return nil
		}
	}
	return errors.New("Wrong appID")
}

func (p *DumbPersistence) Print() {
	fmt.Println("\nUsers:")
	for _, user := range p.Users {
		fmt.Println("----------------------------")
		fmt.Println(user.ToString())
	}
	fmt.Println("\nApplications:")
	for _, app := range p.Apps {
		fmt.Println("----------------------------")
		fmt.Println(app.ToString())
	}
	fmt.Print("\n\n")
}