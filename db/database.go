package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	if DB == nil {
		var err error
		dcs := os.Getenv("DB_URL")
		DB, err = gorm.Open(postgres.Open(dcs), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		DB, db_err := DB.DB()

		if db_err != nil {
			log.Fatalf("Failed to connect to pool: %v", err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		DB.SetMaxIdleConns(5)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		DB.SetMaxOpenConns(30)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		DB.SetConnMaxLifetime(time.Hour)
	} else {
		log.Println("Using existing database connection.")
	}
}
