package controllers

import (
	"fmt"
	"net/http"

	"code/db"
	"code/domain"

	"github.com/lib/pq"
)

type UserAlreadySignedInError struct {
	user_name string
}

func (err *UserAlreadySignedInError) Error() string {
	return fmt.Sprintf("User: %s is already signed in.", err.user_name)
}

func error_handler(
	response http.ResponseWriter,
	err *pq.Error,
	user_name string,
) {
	// User already signed in
	if err.Code == "23505" {
		http.Error(
			response,
			fmt.Sprintf("User: %s already signed in.", user_name),
			http.StatusBadRequest,
		)
		return
	}
	http.Error(
		response,
		fmt.Sprintf(
			"During signing user: %s an unexpected error "+
				"has been ocurred: %s",
			user_name,
			err,
		),
		http.StatusBadRequest)
}

func SignInController(
	response http.ResponseWriter,
	adapter *db.DBAdapter,
	user *domain.UserData,
) {
	err := adapter.SignIn(user)
	if err != nil {
		error_handler(response, err.(*pq.Error), user.UserName)
		return
	}
	msg := fmt.Sprintf("User: %s signed in successful", user.UserName)
	response.Write([]byte(msg))
}
