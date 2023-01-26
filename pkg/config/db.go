package config

import (
	"database/sql"
	"log"
)

func NewDB(driver string, dns string) *sql.DB {
	db, err := sql.Open(driver, dns)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
