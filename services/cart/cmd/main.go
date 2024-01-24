package main

import (
	"cart/internal/logger"
	"cart/internal/messaging"
	sharedcart "cart/internal/sharedCart"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client

var log *zap.SugaredLogger

func init() {
	godotenv.Load()
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	log = logger.GetLogger()

}

func main() {
	sharedcart.Init(redisClient)
	messaging.SetupHTTPServer()
	select {}
}
