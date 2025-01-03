package db

import (
	"fmt"
	"github.com/alpgozbasi/dms-ocr/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func ConnectPostgres(cfg config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to PostgreSQL successfully")

	// i want to migrate manually :)
	err = createDocumentsTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createDocumentsTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS documents (
		id SERIAL PRIMARY KEY,
		file_name VARCHAR(255) NOT NULL,
		file_path VARCHAR(500) NOT NULL,
		ocr_text TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	log.Println("Ensured documents table exists")
	return nil
}
