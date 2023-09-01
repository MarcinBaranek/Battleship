package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code/controllers"
	"code/db"
	"code/domain"
	"code/env"
	"code/middleware"
)

func main() {
	application := http.NewServeMux()

	application.Handle("/", http.HandlerFunc(HelloServer))
	http.Handle(
		"/sign_in",
		middleware.ValidatePostRequest(
			middleware.LogRequest(
				http.HandlerFunc(HandleSignIn),
			),
		),
	)
	http.ListenAndServe(env.Port, nil)
}

func HelloServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])
}

func HandleSignIn(response http.ResponseWriter, request *http.Request) {
	var user domain.UserData
	db_adapter := db.NewDBAdapter()

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(response, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	controllers.SignInController(response, db_adapter, &user)
}
