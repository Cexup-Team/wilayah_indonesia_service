package app

import (
	"database/sql"
	"time"
	"wilayah_indonesia_service/helper"

	_ "github.com/jackc/pgx/v4/stdlib"
)

/**
 *  Todo: connect to database
 *
 *  *please change this credentials connection
 */
func NewDB() *sql.DB {
	// db, err := sql.Open("mysql", "all:password@tcp(0.0.0.0:6603)/db_wilayah")
	db, err := sql.Open("pgx", "postgres://dbmasteruser:postgres123!@$@43.231.128.35:5432/db_wilayah")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
