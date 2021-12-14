package app

import (
	"api-test/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**
 *  Todo: connect to database
 *
 *  *please change this credentials connection
 */
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "all:password@tcp(0.0.0.0:6603)/db_wilayah")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
