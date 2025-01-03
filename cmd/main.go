package main

import (
	"fmt"
	"github.com/alpgozbasi/dms-ocr/config"
	"github.com/alpgozbasi/dms-ocr/internal/api"
	"github.com/alpgozbasi/dms-ocr/internal/db"
	"log"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// connect to the db
	database, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("cannot connect to PostgreSQL: %v", err)
	}
	defer database.Close()

	// initialize router
	router := api.SetupRouter(database)

	// start the server
	serverPort := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("server is running on %s", serverPort)

	err = router.Run(serverPort)
	if err != nil {
		log.Fatalf("cannot run server: %v", err)
	}
}
