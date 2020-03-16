package model

type User struct{
	UserInfo string
	Plan Plan //enum
	Apps []App
}