package model

type PublicApp struct {
	appId   string
	appName string
	limits  Limit
	user    *User
}

func NewPublicApp(appName string, user *User) *PublicApp {
	return &PublicApp{
		appName: appName,
		limits:  Limit{}, //CONFIG
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
	panic("implement me")
}

func (p *PublicApp) GetInfo() string {
	return p.appName
}

func (p *PublicApp) GetUser() *User {
	return p.user
}
