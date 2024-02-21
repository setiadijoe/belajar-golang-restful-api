package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"programmerzamannow/belajar-golang-dependencies-injection/helper"
)

func NewDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "belajar_database_golang")
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
