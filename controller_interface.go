package br_engineer_task

type controller interface {
	CreateUser(userName string, plan string)
	CreateApp(userID string, isOpenSource string)
	ChangeLimits(appID string, concBuild int, buildTime int, buildPerMonth int, teamMembers int)
	OptOutLimits(appID string)
	UsePublicLimits(appID string)
	GetLimit(appID string)
}
