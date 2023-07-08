package storage

import (
	"time"
)

// DATABASE OBJECTS

type User struct {
	Id           string    `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"-"`
	DateCreated  time.Time `db:"date_created" json:"date_created"`
	DateUpdated  time.Time `db:"date_updated" json:"date_updated"`
}

// TABLE DATA

var userTable Table = Table{
	Name: "users",
	CreationQuery: `
    CREATE TABLE users (
      id uuid DEFAULT uuid_generate_v4(),
      username VARCHAR NOT NULL,
      password_hash VARCHAR NOT NULL,
      date_created DATETIME NOT NULL DEFAULT NOW(),
      date_updated DATETIME NOT NULL DEFAULT NOW()
    );
  `,
}

var tables = [...]Table{userTable}

/*
Represents a table in SQL. Contains data for migration as well.
*/
type Table struct {
	Name          string
	CreationQuery string
}
