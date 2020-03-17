package model

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

func (p PrivateApp) SetId(id string) {
	p.appId = id
}

func (p PrivateApp) GetLimits() Limit {
	panic("implement me")
}

func (p PrivateApp) GetInfo() string {
	panic("implement me")
}

func (p PrivateApp) GetUser() *User {
	return p.user
}

