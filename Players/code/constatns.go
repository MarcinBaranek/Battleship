package main

import (
	"fmt"
	"os"
)

func getenv(key, default_ string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return default_
	}
	return value
}

var Port string = fmt.Sprintf(":%s", getenv("PLAYERS_PORT", "8080"))
