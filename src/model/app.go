package model


type App interface {
	GetLimits() *Limit
	SetLimit(limit Limit) error
	GetInfo() string
	SetId(id string)
	GetId() string
	GetUser() *User
	ToString() string
}
