package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort       int
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

func LoadConfig() (Config, error) {
	portStr := os.Getenv("SERVER_PORT")
	if portStr == "" {
		portStr = "8080"
	}
	serverPort, err := strconv.Atoi(portStr)
	if err != nil {
		return Config{}, err
	}

	pgHost := os.Getenv("POSTGRES_HOST")
	if pgHost == "" {
		pgHost = "localhost"
	}

	pgPortStr := os.Getenv("POSTGRES_PORT")
	if pgPortStr == "" {
		pgPortStr = "5432"
	}
	pgPort, err := strconv.Atoi(pgPortStr)
	if err != nil {
		return Config{}, err
	}

	pgUser := os.Getenv("POSTGRES_USER")
	if pgUser == "" {
		pgUser = "postgres"
	}

	pgPassword := os.Getenv("POSTGRES_PASSWORD")

	pgDB := os.Getenv("POSTGRES_DB")
	if pgDB == "" {
		pgDB = "postgres"
	}

	cfg := Config{
		ServerPort:       serverPort,
		PostgresHost:     pgHost,
		PostgresPort:     pgPort,
		PostgresUser:     pgUser,
		PostgresPassword: pgPassword,
		PostgresDB:       pgDB,
	}
	return cfg, nil
}
