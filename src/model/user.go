package model

import "fmt"

type User struct{
	userId   string
	userName string
	plan     Plan
	apps     []*App
}

func NewUser(userName string, plan Plan) User {
	return User{
		userName: userName,
		plan:     plan,
		apps:     nil,
	}
}

func (user *User) GetUserName() string {
	return user.userName
}

func (user *User) AddApp(app App) {
	user.apps = append(user.apps, &app)
}

func (user *User) GetApps() []*App {
	return user.apps
}

func (user *User) SetPlan(plan Plan) {
	user.plan = plan
}

func (user *User) GetPlan() Plan {
	return user.plan
}

func (user *User) SetId(id string) {
	user.userId = id
}

func (user *User) GetId() string {
	return user.userId
}

func (user *User) ToString() string {
	str :=fmt.Sprintf("id: %v\nname: %v\nplan: %v\napps:", user.GetId(), user.GetUserName(), user.GetPlan().Name)
	for _, app:= range user.GetApps() {
		a := *app
		str = str + fmt.Sprintf("   %v", a.GetId())
	}
	return str
}