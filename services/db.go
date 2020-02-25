package services

import (
	"fmt"
	"os"

	"github.com/Davd96/graphql-go-todo/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type connection struct {
	db *sqlx.DB
}

var Connection *connection

func GetConnection() *sqlx.DB {
	if Connection == nil {
		conn, err := sqlx.Connect("postgres", os.Getenv(configs.ENVIROMENT_VARIABLE_PSQL))
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		Connection = &connection{
			db: conn,
		}
	}
	return Connection.db
}
