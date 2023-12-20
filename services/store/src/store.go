package main

import (
	"database/sql"
	"encoding/json"
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

	nc.Subscribe("addToStock", func(m *nats.Msg) {
		var stock database.Stock
		if err := json.Unmarshal(m.Data, &stock); err != nil {
			log.Println("Error al deserializar mensaje:", err)
			nc.Publish(m.Reply, []byte("Error en los datos de registro"))
			return
		}

		ok, err := database.AddToStock(db, stock)
		if err != nil || !ok {
			log.Println("Error al registrar stock:", err)
			nc.Publish(m.Reply, []byte("Error al registrar stock"))
			return
		}

		nc.Publish(m.Reply, []byte("Stock registrado"))
	})

	nc.Subscribe("getStock", func(m *nats.Msg) {
		var product database.Stock
		if err := json.Unmarshal(m.Data, &product); err != nil {
			log.Println("Error al deserializar mensaje:", err)
			nc.Publish(m.Reply, []byte("Error en los datos de registro"))
			return
		}

		stock, err := database.GetStock(db, product.Product)
		if err != nil {
			log.Println("Error al obtener stock:", err)
			nc.Publish(m.Reply, []byte("Error al obtener stock"))
			return
		}

		response, err := json.Marshal(stock)
		if err != nil {
			log.Println("Error al serializar respuesta:", err)
			nc.Publish(m.Reply, []byte("Error al obtener stock"))
			return
		}

		nc.Publish(m.Reply, response)
	})

	nc.Subscribe("removeId", func(m *nats.Msg) {
		var id string
		if err := json.Unmarshal(m.Data, &id); err != nil {
			log.Println("Error al deserializar mensaje:", err)
			nc.Publish(m.Reply, []byte("Error en los datos de registro"))
			return
		}

		ok, err := database.RemoveId(db, id)
		if err != nil || !ok {
			log.Println("Error al registrar stock:", err)
			nc.Publish(m.Reply, []byte("Error al registrar stock"))
			return
		}

		nc.Publish(m.Reply, []byte("Stock eliminado"))
	})

	log.Println("Authentication service is running and listening on NATS")
	select {}
}
