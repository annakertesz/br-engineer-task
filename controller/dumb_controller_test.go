package controller

import (
	"github.com/annakertesz/br-engineer-task/model"
	"github.com/annakertesz/br-engineer-task/persistence"
	"github.com/stretchr/testify/assert"
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
