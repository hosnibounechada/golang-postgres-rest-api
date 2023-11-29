package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDatabase(connection string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	DB = db
	// Perform any database configuration, such as setting up connection pool options.

	return db, nil
}
