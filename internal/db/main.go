package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	gorm.DB
}

func Connect() (*DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_URL")), &gorm.Config{})

	if err != nil {
		return &DB{}, err
	}

	err = db.AutoMigrate(&Host{}, &Repository{})

	if err != nil {
		return &DB{}, err
	}

	// Add hosts to the database
	if err := db.Where("name = ?", "github").First(&Host{}).Error; err != nil {
		db.Create(&Host{
			Name:   "github",
			URL:    "https://github.com/",
			Prefix: "github.com",
		})
	}

	return &DB{*db}, nil
}
