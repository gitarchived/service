package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	gorm.DB
}

func Connect() (DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_URL")), &gorm.Config{})

	db.AutoMigrate(&Host{}, &Repository{})

	if err != nil {
		return DB{}, err
	}

	return DB{DB: *db}, nil
}
