package db

import (
	"database/sql"
	"fmt"
	"log"

	"code/domain"
	"code/env"

	_ "github.com/lib/pq"
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
		log.Panic("Cant connect to the data base")
		panic(err)
	}
	if err = conn.Ping(); err != nil {
		log.Panic("Data base is not responding on ping")
		panic(err)
	}
	return &DBAdapter{
		connection: conn,
	}
}

func (adapter *DBAdapter) SignIn(user_data *domain.UserData) error {
	query := fmt.Sprintf(
		"INSERT INTO public.user (user_name, password_hash) VALUES ('%s', '%s')",
		user_data.UserName, user_data.PasswordHash,
	)
	_, err := adapter.connection.Exec(query)
	return err
}
