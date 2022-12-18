package database

import (
	"database/sql"
	"github.com/pressly/goose"
	"time"
)

type Db struct {
	*sql.DB
}

func NewDB(dsn string, automigrate bool) (*Db, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Minute)

	if automigrate {
		err := goose.SetDialect("sqlite3")
		if err != nil {
			return nil, err
		}

		err = goose.Up(db, "./assets/mirgations")
		if err != nil {
			return nil, err
		}
	}
	return &Db{db}, nil
}
