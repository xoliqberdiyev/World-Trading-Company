package config

import (
	"log"

	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Databasse string
}

func Load() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("error to load .env file")
	}
	return &Config{
		Postgres: Postgres{
			User:      utils.GetString("DB_USER", "postgres"),
			Password:  utils.GetString("DB_PASSWORD", "20090912"),
			Host:      utils.GetString("DB_HOST", "localhost"),
			Port:      utils.GetString("DB_PORT", "5432"),
			Databasse: utils.GetString("DB_NAME", "wtc_db"),
		},
	}
}
