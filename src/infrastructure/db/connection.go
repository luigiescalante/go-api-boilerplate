package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go.api-boilerplate/config"
	"log"
)

var db *sql.DB

func Db() (*sql.DB, error) {
	env := config.GetConfig()
	if db == nil {
		conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.GetDbHost(),
			env.GetDbPort(),
			env.GetDbUser(),
			env.GetDbPassword(),
			env.GetDbName())
		db, err := sql.Open("postgres", conn)
		if err != nil {
			log.Panicf("error on db conn:%s", err.Error())
		}
		return db, nil
	} else {
		return db, nil
	}
}
