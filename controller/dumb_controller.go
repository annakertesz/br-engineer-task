package controller

import (
	br_engineer_task "github.com/annakertesz/br-engineer-task"
	"github.com/annakertesz/br-engineer-task/model"
)

type DumbController struct {
	db br_engineer_task.Persist
	plans model.PlanType
}

func NewDumbController(db br_engineer_task.Persist, plans model.PlanType) DumbController {
	return DumbController{
		db:    db,
		plans: plans,
	}
}

func (d DumbController) CreateUser(userName string, planString string) model.User {
	user := model.NewUser(userName, d.plans.Get(planString))
	d.db.Save(&user)
	return user
}

func (d DumbController) CreateApp(userID string, isOpenSource string) {
	panic("implement me")
}

func (d DumbController) ChangeLimits(appID string, concBuild int, buildTime int, buildPerMonth int, teamMembers int) {
	panic("implement me")
}

func (d DumbController) OptOutLimits(appID string) {
	panic("implement me")
}

func (d DumbController) UsePublicLimits(appID string) {
	panic("implement me")
}

func (d DumbController) GetLimit(appID string) {
	panic("implement me")
}

