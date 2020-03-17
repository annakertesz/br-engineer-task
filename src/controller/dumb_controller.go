package controller

import (
	"errors"
	"github.com/annakertesz/br-engineer-task/src/config"
	"github.com/annakertesz/br-engineer-task/src/model"
	"github.com/annakertesz/br-engineer-task/src/persistence"
	"time"
)

type DumbController struct {
	db                persistence.Persist
	plans             model.PlanType
	opensourceDefault model.Limit
}

func NewDumbController(db persistence.Persist, config config.Config) DumbController {
	return DumbController{
		db:    db,
		plans: config.Plans,
		opensourceDefault:config.OpensourceDefault,
	}
}

func (d DumbController) CreateUser(userName string, planString string) model.User {
	user := model.NewUser(userName, d.plans.Get(planString))
	d.db.SaveUser(&user)
	return user
}

func (d DumbController) CreateApp(userID string, appName string, openSource bool) (model.App, error){
	user := d.db.GetUser(userID)
	if user == nil {
		return nil, errors.New("Wrong userID")
	}
	var newApp model.App
	if openSource {
		newApp = model.NewPublicApp(appName, user, d.opensourceDefault)
	} else {
		newApp = model.NewPrivateApp(appName, user)
	}
	user.AddApp(newApp)
	d.db.SaveApp(newApp)
	d.db.UpdateUser(*user)
	return newApp, nil
}

func (d DumbController) ChangeLimits(appID string, concBuild int, buildTime time.Duration, buildPerMonth int, teamMembers int) error {
	app := d.db.GetApp(appID)
	err := app.SetLimit(model.Limit{
		ConcurrentBuild: concBuild,
		BuildTime:       model.Duration{buildTime},
		BuildsPerMonth:  buildPerMonth,
		TeamMembers:     teamMembers,
	})
	if err != nil {
		return err
	}
	d.db.UpdateApp(app)
	return nil
}

func (d DumbController) UsePrivateLimits(appID string) error {
	app := d.db.GetApp(appID)
	publicApp, ok := app.(*model.PublicApp)
	if !ok {
		return errors.New("this is already a private application")
	}
	privateApp := publicApp.TransformToPrivate()
	d.db.UpdateApp(privateApp)
	return nil
}

func (d DumbController) GetLimit(appID string) model.Limit {
	app := d.db.GetApp(appID)
	return app.GetLimits()
}

