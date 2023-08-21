package middleware

import (
	"log"
	"net/http"
)

func ValidatePostRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			if request.Method != http.MethodPost {
				http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			if request.ContentLength == 0 {
				http.Error(response, "Got request with empty body!", http.StatusBadRequest)
				return
			}
			if request.Header.Get("Content-Type") != "application/json" {
				http.Error(response, "Expected content-type application/json", http.StatusBadRequest)
				return
			}
			log.Println("Request validation: OK")

			next.ServeHTTP(response, request)
		},
	)
}
