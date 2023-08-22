package db

import (
	"database/sql"
	"fmt"
	"log"

	"code/domain"
	"code/env"
)

type DBAdapter struct {
	connection *sql.DB
}

func (adapter *DBAdapter) Close() {
	adapter.connection.Close()
}

func NewDBAdapter() *DBAdapter {
	conn, err := sql.Open("postgres", env.DB_URL)
	if err != nil {
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	}
	return &DBAdapter{
		connection: conn,
	}
}

func (adapter *DBAdapter) SignIn(user_data *domain.UserData) error {
	query := fmt.Sprintf(
		"INSERT INTO users VALUES(user_name, password_hash) (%s, %s)",
		user_data.UserName, user_data.PasswordHash,
	)
	_, err := adapter.connection.Exec(query)
	if err != nil {
		log.Fatalf(
			"During inserting into data base user: %s, got the error: %s",
			user_data.UserName, err,
		)
		return err
	}
	return nil
}
