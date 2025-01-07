package config

import "os"

type db struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSL      string
}

func loadDatabaseConfig() db {
	return db{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSL:      os.Getenv("DB_SSL_MODE"),
	}
}

func loadDatabaseTestConfig() db {
	return db{
		Host:     os.Getenv("TEST_DB_HOST"),
		Port:     os.Getenv("TEST_DB_PORT"),
		User:     os.Getenv("TEST_DB_USER"),
		Password: os.Getenv("TEST_DB_PASSWORD"),
		Database: os.Getenv("TEST_DB_NAME"),
		SSL:      os.Getenv("TEST_DB_SSL_MODE"),
	}
}
