package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

func NewDB() (*DB, error) {
	readDBConnect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8&timeout=%s&tls=%s",
		os.Getenv("APP_MYSQL_USER"),
		os.Getenv("READ_DB_PASSWORD"),
		os.Getenv("READ_DB_HOST"),
		os.Getenv("DB_NAME"),
		"30s",
		"false",
	)

	writeDBConnect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8&timeout=%s&tls=%s",
		os.Getenv("APP_MYSQL_USER"),
		os.Getenv("WRITE_DB_PASSWORD"),
		os.Getenv("WRITE_DB_HOST"),
		os.Getenv("DB_NAME"),
		"30s",
		"false",
	)

	readDB, err := sqlx.Open("mysql", readDBConnect)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	writeDB, err := sqlx.Open("mysql", writeDBConnect)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	return &DB{
		Read:  readDB,
		Write: writeDB,
	}, nil
}

func (db *DB) Close() {
	_ = db.Read.Close()
	_ = db.Write.Close()
}
