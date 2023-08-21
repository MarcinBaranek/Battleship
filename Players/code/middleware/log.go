package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			bodyBytes, _ := io.ReadAll(request.Body)
			// Wrap it back for further processing
			request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			log.Println("Raw Request Body:", string(bodyBytes))

			next.ServeHTTP(response, request)
		},
	)
}
