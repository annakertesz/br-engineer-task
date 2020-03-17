package model

type App interface {
	GetLimits() Limit
	GetInfo() string
	SetId(id string)
	GetUser() *User
}
