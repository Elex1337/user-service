package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"log"
)

func RunMigrations(db *sqlx.DB, dir string) {
	sqlDB := db.DB

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	if err := goose.Up(sqlDB, dir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}
