package model

import "errors"

type PrivateApp struct {
	appId string
	appName string
	isOpenSource bool
	user *User
}

func NewPrivateApp(appName string, user *User) *PrivateApp {
	return &PrivateApp{
		appName:appName,
		isOpenSource:false,
		user:user,
	}
}

func (p *PrivateApp) SetId(id string) {
	p.appId = id
}

func (p *PrivateApp) GetId() string {
	return p.appId
}

func (p *PrivateApp) GetLimits() Limit {
	return p.user.plan.Limits
}

func (p *PrivateApp) SetLimit(limit Limit) error {
	return errors.New("Cant change limits of private application")
}

func (p *PrivateApp) GetInfo() string {
	return p.appName
}

func (p PrivateApp) GetUser() *User {
	return p.user
}

