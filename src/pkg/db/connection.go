package db

import (
	configs "api-project/src/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name)

	conn, err := sql.Open("postgres", stringConnection)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}