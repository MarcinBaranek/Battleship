package main

import (
	"bytes"
	"database/sql"
	"io"
	"net/http"
	"testing"

	"code/env"
)

func SetUp() {
	conn, err := sql.Open("postgres", env.DB_URL)
	if err != nil {
		panic(err)
	}
	conn.Exec("DELETE FROM public.user")
	conn.Close()
}

func SendBody(body []byte, t *testing.T) string {
	var response *http.Response
	var parsed_response []byte
	bodyReader := bytes.NewReader(body)
	response, err := http.Post(
		"http://localhost:8080/sign_in", "application/json", bodyReader,
	)
	if err != nil {
		t.Fatalf("During post request got the error: %s", err)
	}
	defer response.Body.Close()

	parsed_response, err = io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(parsed_response)
}

func TestSignIn(t *testing.T) {
	SetUp()
	defer SetUp()

	var parsed_response []byte

	jsonBody := []byte(`{"UserName":"John","PasswordHash":"RandomHash"}`)
	response := SendBody(jsonBody, t)

	exp_msg := "User: John signed in successful"
	if response != exp_msg {
		t.Fatalf(
			"got:\n%s!=\n%s", string(parsed_response), exp_msg,
		)
	}

	jsonBody = []byte(`{"UserName":"John","PasswordHash":"RandomHash"}`)
	response = SendBody(jsonBody, t)

	exp_msg = "User: John already signed in.\n"
	if response != exp_msg {
		t.Fatalf(
			"got:\n%s!=\n%s", string(parsed_response), exp_msg,
		)
	}
}
