package main

import (
	"core/internal/logger"
	"core/internal/messaging"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	log = logger.GetLogger()

}

func main() {
	messaging.SetupHTTPServer()
	select {}
}
