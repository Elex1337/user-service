package db

import (
	"fmt"
	"github.com/Elex1337/user-service/config"
	"github.com/jmoiron/sqlx"
	"log"
)

func ConnectDB(cfg *config.DBConfig) *sqlx.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")

	return db
}
