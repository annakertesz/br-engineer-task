package persistence

import (
	"github.com/annakertesz/br-engineer-task/model"
	"github.com/lithammer/shortuuid"
)

type DumbPersistence struct{
	users []*model.User
}

func (p *DumbPersistence) Save(user *model.User) {
	user.SetId(shortuuid.New())
	p.users = append(p.users, user)
}

