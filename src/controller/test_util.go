package controller

import (
	"github.com/annakertesz/br-engineer-task/src/config"
	"github.com/annakertesz/br-engineer-task/src/model"
	"github.com/annakertesz/br-engineer-task/src/persistence"
	)

const (
	USER_NAME_A="usernameA"
	USER_ID_A="userIdA"
	USER_PLAN_A="free"
	USER_NAME_B="usernameB"
	USER_ID_B="userIdB"
	USER_PLAN_B="developer"

	PRIVATE_APP_NAME_A="appnameA"
	PRIVATE_APP_ID_A="appIDA"
	PUBLIC_APP_NAME_B="appnameB"
	PUBLIC_APP_ID_B="appIDB"
)


func GetControllerWithEmptyPersistence() DumbController {
	p := persistence.NewDumbPersistence()
	config, err := config.GetConfigFromFile("../config/limit_config.json") //TODO: should read config struct here
	if err != nil {
		panic("couldnt read config")
	}
	return NewDumbController(p, *config)
}

func GetControllerWithDataInPersistence() DumbController {
	config, err := config.GetConfigFromFile("../config/limit_config.json") //TODO: should read config struct here
	if err != nil {
		panic("couldnt read config")
	}
	p := getPersistenceWithData(*config)
	return NewDumbController(&p, *config)
}

func getPersistenceWithData(config config.Config) persistence.DumbPersistence {
	userA := getExampleUser(USER_NAME_A, USER_ID_A, config.Plans.Get(USER_PLAN_A))
	userB := getExampleUser(USER_NAME_B, USER_ID_B, config.Plans.Get(USER_PLAN_B))
	p := persistence.DumbPersistence{
		Users: []*model.User{
			userA,
			userB,
		},
		Apps: []model.App{
			getExamplePrivateApp(PRIVATE_APP_NAME_A, PRIVATE_APP_ID_A, userA),
			getExamplePublicApp(PUBLIC_APP_NAME_B, PUBLIC_APP_ID_B, userA, config.OpensourceDefault),
		},
	}
	return p
}

func getExampleUser(name string, id string, plan model.Plan) *model.User {
	user := model.NewUser(
		name,
		plan)
	user.SetId(id)
	return &user
}

func getExamplePublicApp(name string, id string, user *model.User, limit model.Limit) model.App {
	app := model.NewPublicApp(name, user, limit)
	app.SetId(id)
	return app
}

func getExamplePrivateApp(name string, id string, user *model.User) model.App {
	app := model.NewPrivateApp(name, user)
	app.SetId(id)
	return app
}
