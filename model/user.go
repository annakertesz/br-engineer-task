package model

type User struct{
	userInfo string
	plan Plan
	apps []App
}

func NewUser(userInfo string, plan Plan) User {
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

func (user *User) SetPlan(plan Plan) {
	user.plan = plan
}

func (user *User) GetPlan() Plan {
	return user.plan
}