package br_engineer_task

import "github.com/annakertesz/br-engineer-task/model"

type Persist interface {
	Save(user *model.User) string
}
