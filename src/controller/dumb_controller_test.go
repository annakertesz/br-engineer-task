package controller

import (
	"github.com/annakertesz/br-engineer-task/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const PRIVATE_APP_TO_CREATE_NAME  = "new private app"
const PUBLIC_APP_TO_CREATE_NAME  = "new public app"

func TestDumbController_CreateUser(t *testing.T) {
	c := GetControllerWithEmptyPersistence()
	user := c.CreateUser("UserName", "free")
	assert.Equal(t, 1, len(c.db.GetUsers()))
	assert.Equal(t, "UserName", user.GetUserName())
	assert.Equal(t, time.Duration(10*time.Minute), user.GetPlan().Limits.BuildTime.Duration)
}

func TestDumbController_CreateApp(t *testing.T) {
	c := GetControllerWithDataInPersistence()
	_, err := c.CreateApp(USER_ID_B, PRIVATE_APP_TO_CREATE_NAME, false)
	require.NoError(t, err)
	user := c.db.GetUser(USER_ID_B)
	require.Equal(t, 1, len(user.GetApps()))
	persistedApp := *user.GetApps()[0]
	assert.Equal(t, persistedApp.GetUser(), user)
	assert.IsType(t, &model.PrivateApp{}, persistedApp)

	//create public app
	c.CreateApp(USER_ID_B, PUBLIC_APP_TO_CREATE_NAME, true)
	user = c.db.GetUser(USER_ID_B)
	require.Equal(t, 2, len(user.GetApps()))
	persistedApp = *user.GetApps()[1]
	assert.Equal(t, persistedApp.GetUser(), user)
	assert.IsType(t, &model.PublicApp{}, persistedApp)

}
func TestDumbController_GetLimit(t *testing.T) {
	c := GetControllerWithDataInPersistence()
	assert.Equal(t, 2, c.GetLimit(PUBLIC_APP_ID_B).ConcurrentBuild)
	assert.Equal(t, 1, c.GetLimit(PRIVATE_APP_ID_A).ConcurrentBuild)

}

func TestDumbController_ChangeLimits(t *testing.T) {
	c := GetControllerWithDataInPersistence()
	err := c.ChangeLimits(PUBLIC_APP_ID_B, 3, 3, 3, 3)
	assert.NoError(t, err)
	assert.Equal(t, 3, c.GetLimit(PUBLIC_APP_ID_B).ConcurrentBuild)
	err = c.ChangeLimits(PRIVATE_APP_ID_A, 3, 3, 3, 3)
	assert.Error(t, err)
	assert.Equal(t, 1, c.GetLimit(PRIVATE_APP_ID_A).ConcurrentBuild)
}

func TestDumbController_UsePrivateLimits(t *testing.T) {
	c := GetControllerWithDataInPersistence()
	err := c.UsePrivateLimits(PUBLIC_APP_ID_B)
	assert.NoError(t, err)
	app := c.db.GetApp(PUBLIC_APP_ID_B)
	assert.IsType(t, &model.PrivateApp{}, app)
	err = c.UsePrivateLimits(PUBLIC_APP_ID_B)
	assert.Error(t, err)
}



