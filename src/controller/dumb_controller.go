package controller

import (
	"errors"
	"github.com/annakertesz/br-engineer-task/src/config"
	"github.com/annakertesz/br-engineer-task/src/model"
	"github.com/annakertesz/br-engineer-task/src/persistence"
	"time"
)

type DumbController struct {
	db                persistence.Persistence
	plans             model.PlanType
	opensourceDefault model.Limit
}

func NewDumbController(db persistence.Persistence, config config.Config) *DumbController {
	return &DumbController{
		db:    db,
		plans: config.Plans,
		opensourceDefault:config.OpensourceDefault,
	}
}

func (d *DumbController) CreateUser(userName string, planString string) (*model.User, error) {
	plan, err := d.plans.Get(planString)
	if err !=nil {
		return nil, err
	}
	user := model.NewUser(userName, plan)
	d.db.SaveUser(&user)
	return &user, nil
}

func (d *DumbController) CreateApp(userID string, appName string, openSource bool) (model.App, error){
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
	err := d.db.UpdateUser(*user)
	if err != nil {
		return nil, err
	}
	return newApp, nil
}

func (d *DumbController) ChangeLimits(appID string, concBuild int, buildTime time.Duration, buildPerMonth int, teamMembers int) error {
	app := d.db.GetApp(appID)
	if app == nil {
		return errors.New("Wrong userID")
	}
	err := app.SetLimit(model.Limit{
		ConcurrentBuild: concBuild,
		BuildTime:       model.Duration{buildTime},
		BuildsPerMonth:  buildPerMonth,
		TeamMembers:     teamMembers,
	})
	if err != nil {
		return err
	}
	err = d.db.UpdateApp(app)
	if err != nil {
		return err
	}
	return nil
}

func (d *DumbController) UsePrivateLimits(appID string) error {
	app := d.db.GetApp(appID)
	if app == nil {
		return errors.New("Wrong userID")
	}
	publicApp, ok := app.(*model.PublicApp)
	if !ok {
		return errors.New("this is already a private application")
	}
	privateApp := publicApp.TransformToPrivate()
	err := d.db.UpdateApp(privateApp)
	if err != nil {
		return err
	}
	return nil
}

func (d *DumbController) GetLimit(appID string) (*model.Limit, error) {
	app := d.db.GetApp(appID)
	if app==nil {
		return nil, errors.New("Wrong appID")
	}
	return app.GetLimits(), nil
}

