package main

import (
	"fmt"
	"os"

	"auth/internal/logger"
	"auth/internal/messaging"
	"auth/internal/token"

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

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})
	log.Debug("Redis connected")
}

func main() {
	token.Init(redisClient)
	messaging.SetupHTTPServer()

	log.Infof("Auth service started on port %s", os.Getenv("HTTP_PORT"))
	select {}
}
