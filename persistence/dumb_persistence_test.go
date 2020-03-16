package persistence

import (
	"github.com/annakertesz/br-engineer-task/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDumbPersistence_GetUser(t *testing.T) {
	p := NewDumbPersistence()
	//TODO: put into util
	userA := model.NewUser(
		"User One",
		model.Plan{
			Name:   "free",
			Price:  0,
			Limits: model.Limit{1, model.Duration{time.Minute}, 1, 1},
		})
	userA.SetId("idA")
	userB := model.NewUser(
		"User Two",
		model.Plan{
			Name:   "dev",
			Price:  0,
			Limits: model.Limit{2, model.Duration{time.Hour}, 2, 2},
		})
	userB.SetId("idB")
	p.users = []*model.User{&userA,&userB}
	user := p.GetUser("idA")
	assert.Equal(t, "User One", user.GetUserName())
	user = p.GetUser("idC")
	assert.Nil(t, user)
}
