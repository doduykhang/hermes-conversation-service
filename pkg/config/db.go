package config

import (
	"database/sql"
	"fmt"
	"log"
)

func NewDB(config *Config) *sql.DB {
	db, err := sql.Open(config.DB.Driver, getDBConnString(config.DB))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getDBConnString(db DB) string {
	//"sammy:password@/hermes_conversation?parseTime=True&loc=Local",
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
