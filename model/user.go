package model

type User struct{
	userInfo string
	plan Limit //TODO: Plan should be plan
	apps []App
}

func NewUser(userInfo string, plan Limit) User {
	return User{
		userInfo: userInfo,
		plan:     plan,
		apps:     nil,
	}
}

func (user *User) AddApp(app App) {
	user.apps = append(user.apps, app)
}

func (user *User) GetApps() []App {
	return user.apps
}

func (user *User) SetPlan(limit Limit) {
	user.plan = limit
}

func (user *User) GetPlan() Limit {
	return user.plan
}