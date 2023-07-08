package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/*
Database connection to PostgreSQL
*/
var PgDB *sqlx.DB

/*
Creates an instance of SQLx.DB that is accessible under the name PgDB
*/
func Init() {
	// get environment variables
	user := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	pw := os.Getenv("POSTGRES_PASSWORD")

	// establish connection
	connection, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbName, pw))
	if err != nil {
		log.Fatalln(err)
	}
	PgDB = connection

	// migrate
	MigrateIfNeeded()
}

/*
Adds missing tables to the database
*/
func MigrateIfNeeded() {
	// if there are no tables, apply migrations
	for i, table := range tables {
		log.Printf("Checking table [%d] %s\n", i, table.Name)
		res, err := PgDB.Exec("SELECT EXISTS ( SELECT FROM ?)", table.Name)
		if err != nil {
			return
		}
		fmt.Println(res)
	}
}
