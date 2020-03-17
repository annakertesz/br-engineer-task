package model

import (
	"time"
)

type PublicApp struct {
	appId   string
	appName string
	limits  Limit
	user    *User
}

func NewPublicApp(appName string, user *User) *PublicApp {
	return &PublicApp{
		appName: appName,
		limits:  Limit{  //TODO: from config
			ConcurrentBuild: 2,
			BuildTime:       Duration{time.Minute*45},
			BuildsPerMonth:  -1,
			TeamMembers:     -1,
		}, //CONFIG
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
