package main

import (
	"database/internal/database"
	"database/internal/logger"
	"database/internal/messaging"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func main() {

	db, err := database.Init()
	if err != nil {
		log.Fatal("Error initializing database")
	}
	defer db.Close()

	messaging.SetupHttp(db)

	log.Infof("Database service started")
	select {}
}
