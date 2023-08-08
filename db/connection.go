package db

import (
	"database/sql"
	"example/configs"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	dbConfiguration := configs.GetDB()

	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfiguration.Host, dbConfiguration.Port, dbConfiguration.User, dbConfiguration.Pass, dbConfiguration.Database,
	)

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
