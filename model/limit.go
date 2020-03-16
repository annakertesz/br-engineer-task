package model

import "time"

type Limit struct {
	ConcurrentBuild int
	BuildTime time.Duration
	BuildsPerMonth int
	TeamMembers int
}
