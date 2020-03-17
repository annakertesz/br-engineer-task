package model

import "fmt"

type PublicApp struct {
	appId   string
	appName string
	limits  Limit
	user    *User
}

func NewPublicApp(appName string, user *User, limit Limit) *PublicApp {
	return &PublicApp{
		appName: appName,
		limits:  limit,
		user:    user,
	}
}

func (p *PublicApp) SetId(id string) {
	p.appId = id
}
func (p *PublicApp) GetId() string {
	return p.appId
}

func (p *PublicApp) GetLimits() Limit {
	return p.limits
}

func (p *PublicApp) SetLimit(limit Limit) error {
	p.limits = limit
	return nil
}

func (p *PublicApp) GetInfo() string {
	return p.appName
}

func (p *PublicApp) GetUser() *User {
	return p.user
}

func (p *PublicApp) TransformToPrivate() *PrivateApp {
	privateApp := NewPrivateApp(p.appName, p.user)
	privateApp.appId = p.appId
	return privateApp
}

func (p *PublicApp) ToString() string {
	return fmt.Sprintf("type: public\nid: %v\nname:%v\nuser: %v\nlimit: %v",p.GetId(), p.GetInfo(), p.GetUser().GetId(), p.limits)
}
