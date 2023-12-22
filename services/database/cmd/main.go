package main

import (
	"database/internal/database"
	"database/internal/logger"
	"database/internal/messaging"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Init()
	if err != nil {
		log.Fatal("Error initializing database")
	}
	defer db.Close()

	natsConn, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal("Error connecting to NATS")
	}
	defer natsConn.Close()

	messaging.SetupSubscribers(natsConn, db)

	log.Infof("Database service started")
	select {}
}
