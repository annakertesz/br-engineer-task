package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetPlansFromConfig(t *testing.T) {
	planType := GetPlansFromConfig("../config/config.json")
	assert.Equal(t, 2,  planType.Free.Limits.TeamMembers)
	assert.Equal(t, time.Duration(10*time.Minute), planType.Free.Limits.BuildTime.Duration)
}
