package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"store/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var db *sql.DB
var jwtSecret string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	jwtSecret = os.Getenv("JWT_SECRET")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})
	log.Println("Redis connected")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	tempDB, err := database.InitDB(dbURI)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	db = tempDB
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("checkout", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
		nc.Publish(m.Reply, []byte("Payment processed"))
	})

	log.Println("Payment service is running and listening on NATS")
	select {}
}
