package persistence

import (
	"github.com/annakertesz/br-engineer-task/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestDumbPersistence_UpdateUser(t *testing.T) {
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
	userA.AddApp(model.NewPublicApp("app", &userA))  //TODO: dont need userA ref
	p.users = []*model.User{&userA}
	require.Equal(t, "User One", p.users[0].GetUserName())
	require.Equal(t, 1, len(p.users[0].GetApps()))
	userA.AddApp(model.NewPrivateApp("name", &userA))
	p.UpdateUser(userA)
	assert.Equal(t, 2, len(p.users[0].GetApps()))
}
