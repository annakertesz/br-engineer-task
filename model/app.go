package model

type App interface {
	GetLimits() Limit
	GetInfo() string
}
