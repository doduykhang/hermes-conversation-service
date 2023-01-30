package config

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func Migrate(config *Config) {
	m, err := migrate.New(
		"file://db/migration",
		fmt.Sprintf(
			"mysql://%s:%s@tcp(%s:%s)/%s",
			config.DB.User,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Name,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
