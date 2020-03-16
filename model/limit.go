package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Limit struct {
	ConcurrentBuild int `json:"concurrent_build"`
	BuildTime Duration `json:"build_time_min"`
	BuildsPerMonth int `json:"builds_per_month"`
	TeamMembers int `json:"team_members"`
}

type Duration struct {
	time.Duration
}
func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)*time.Minute
		return nil
	case string:
		duration, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		d.Duration = duration*time.Minute
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
