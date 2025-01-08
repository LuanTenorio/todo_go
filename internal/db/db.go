package db

import (
	"fmt"
	"log"

	"github.com/LuanTenorio/todo_go/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Connection *sqlx.DB

func Inicialize() {
	Connection = NewDBConnection(config.Env.DBUrl)
}

func NewDBConnection(dbStrConnection string) *sqlx.DB {
	fmt.Println(dbStrConnection)
	db, err := sqlx.Connect("postgres", dbStrConnection)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return db
}
