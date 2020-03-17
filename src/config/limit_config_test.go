package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetConfigFromFile(t *testing.T) {
	config, err := GetConfigFromFile("../config/limit_config.json")
	require.NoError(t, err)
	assert.Equal(t, 2, config.OpensourceDefault.ConcurrentBuild)
	assert.Equal(t, "organization", config.Plans.Organization.Name)
	assert.Equal(t, time.Minute*10, config.Plans.Free.Limits.BuildTime.Duration)
}
