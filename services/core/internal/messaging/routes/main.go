package routes

import (
	"core/internal/logger"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var authUrl string
var storeUrl string

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	authUrl = os.Getenv("AUTH_SERVICE_URL")
	storeUrl = os.Getenv("STORE_SERVICE_URL")
}
