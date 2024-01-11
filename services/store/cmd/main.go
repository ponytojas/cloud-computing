package main

import (
	"os"
	"store/internal/logger"
	"store/internal/messaging"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client

var log *zap.SugaredLogger

func init() {
	godotenv.Load()
	log = logger.GetLogger()
}

func main() {
	messaging.SetupHTTPServer()

	log.Info("Store service started on port %s", os.Getenv("HTTP_PORT"))
	select {}
}
