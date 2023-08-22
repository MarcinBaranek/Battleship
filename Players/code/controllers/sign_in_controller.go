package controllers

import (
	"code/db"
	"code/domain"
)

func SignInController(adapter *db.DBAdapter, user *domain.UserData) error {
	return adapter.SignIn(user)
}
