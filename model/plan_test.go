package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetPlansFromConfig(t *testing.T) {
	plan := GetPlansFromConfig("../config/config.json")
	assert.Equal(t, 2,  plan.Free.TeamMembers)
	assert.Equal(t, time.Duration(10*time.Minute), plan.Free.BuildTime.Duration)
	assert.Equal(t, -1, plan.OpensourceDefault.TeamMembers)
}
