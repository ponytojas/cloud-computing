package main

import (
	"core/internal/logger"
	"core/internal/messaging"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	godotenv.Load()

	log = logger.GetLogger()

}

func main() {
	messaging.SetupHTTPServer()
	select {}
}
