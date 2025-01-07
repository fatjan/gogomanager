package database

import (
	"fmt"
	"time"
	"log"

	"github.com/fatjan/gogomanager/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitiateDBConnection(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Database, cfg.DB.SSL)

	log.Printf(dsn)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDBConnection(db *sqlx.DB) error {
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
