package main

import (
	"database/internal/database"
	"database/internal/logger"
	"database/internal/messaging"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func Init() {
	log = logger.GetLogger()
}

func main() {
	if os.Getenv("VSCODE_DEBUG") != "true" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db, err := database.Init()
	if err != nil {
		log.Fatal("Error initializing database")
	}
	defer db.Close()

	messaging.SetupHttp(db)

	log.Infof("Database service started")
	select {}
}
