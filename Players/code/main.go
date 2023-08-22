package main

// run with comand: `go run main.go constatns.go`
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	log.Println("start handling sign_in")
	var p domain.UserData

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
		http.Error(response, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("User name: %s\n", p.UserName)
	response.Write([]byte("OK -> Sign in"))
}
