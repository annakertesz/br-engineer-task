package controller

import (
	"github.com/annakertesz/br-engineer-task/model"
	"github.com/annakertesz/br-engineer-task/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDumbController_CreateUser(t *testing.T) {
	p := persistence.NewDumbPersistence()
	planType := model.GetPlansFromConfig("../config/config.json") //TODO: should read config struct here
	c := NewDumbController(p, planType)
	user := c.CreateUser("UserName", "free")
	assert.Equal(t, 1, len(p.GetUsers()))
	assert.Equal(t, "UserName", user.GetUserName())
	assert.Equal(t, time.Duration(10*time.Minute), user.GetPlan().Limits.BuildTime.Duration)
}

func TestDumbController_CreateApp(t *testing.T) {
	p := persistence.NewDumbPersistence()
	planType := model.GetPlansFromConfig("../config/config.json") //TODO: should read config struct here
	c := NewDumbController(p, planType)
	user := c.CreateUser("UserName", "free")
	//Create private app
	c.CreateApp(user.GetId(), "App Name", false)
	users := p.GetUsers()
	require.Equal(t, 1, len(users))
	persistedUser := users[0]
	require.Equal(t, 1, len(persistedUser.GetApps()))
	newPrivateApp := *persistedUser.GetApps()[0]
	privateAppId := newPrivateApp.GetId()
	assert.Equal(t, newPrivateApp.GetUser(), persistedUser)
	assert.IsType(t, &model.PrivateApp{}, newPrivateApp)
	//create public app
	c.CreateApp(user.GetId(), "App Name", true)
	users = p.GetUsers()
	require.Equal(t, 1, len(users))
	persistedUser = users[0]
	require.Equal(t, 2, len(persistedUser.GetApps()))
	newPublicApp := *persistedUser.GetApps()[1]
	publicAppId := newPublicApp.GetId()
	assert.Equal(t, newPublicApp.GetUser(), persistedUser)
	assert.IsType(t, &model.PublicApp{}, newPublicApp)
	//test getLimit function
	assert.Equal(t, 2, c.GetLimit(publicAppId).ConcurrentBuild)
	assert.Equal(t, 1, c.GetLimit(privateAppId).ConcurrentBuild)
	//test changeLimit function
	err := c.ChangeLimits(publicAppId, 3, 3, 3, 3)
	assert.NoError(t, err)
	assert.Equal(t,3,  c.GetLimit(publicAppId).ConcurrentBuild)
	err = c.ChangeLimits(privateAppId, 3, 3, 3, 3)
	assert.Error(t, err)
	assert.Equal(t,1,  c.GetLimit(privateAppId).ConcurrentBuild)

}



