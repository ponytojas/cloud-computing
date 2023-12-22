package main

import (
	"fmt"
	"os"

	"auth/internal/logger"
	"auth/internal/messaging"
	"auth/internal/token"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})
	log.Debug("Redis connected")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	natsConn, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal("Error connecting to NATS")
	}
	defer natsConn.Close()

	token.Init(redisClient)

	messaging.SetupSubscribers(natsConn)

	log.Infof("Database service started")
	select {}
}
