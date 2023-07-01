package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func InitializeDatabaseConnection() *Database {
	session, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Database{session}
}
