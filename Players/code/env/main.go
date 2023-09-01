package env

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

var (
	Port   string = fmt.Sprintf(":%s", getenv("PLAYERS_PORT", "8080"))
	DB_URL string = getenv(
		"DB_URL", "postgres://postgres:password@localhost:5432/battleship_db?sslmode=disable",
	)
)
