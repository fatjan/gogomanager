package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	App    app
	DB     db
	TestDB db
}

func LoadConfig() (*Config, error) {
	envPath, err := getEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		log.Print("LoadEnv() failed, No .env file found, env should be injected by other methods")
		return nil, err
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		App:    loadApplicationConfig(),
		DB:     loadDatabaseConfig(),
		TestDB: loadDatabaseTestConfig(),
	}, nil
}

func getEnvPath() (string, error) {
	directory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filepath := searchup(directory, ".env")
	return filepath, nil
}

func searchup(dir string, filename string) string {
	if dir == "/" || dir == "" || dir == "." {
		return ""
	}

	if _, err := os.Stat(path.Join(dir, filename)); err == nil {
		return path.Join(dir, filename)
	}

	return searchup(path.Dir(dir), filename)
}
