package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"auth/database"

	"github.com/golang-jwt/jwt/v5"
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

func createToken(auth database.AuthCheck) (string, error) {
	ctx := context.Background()
	cachedToken, err := redisClient.Get(ctx, auth.Username).Result()
	// If there's not already a token in Redis it will return an error.
	if err == nil {
		return cachedToken, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   auth.UserId,
		"username": auth.Username,
		"email":    auth.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	err = redisClient.Set(ctx, auth.Username, tokenString, time.Hour*24).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("health", func(m *nats.Msg) {
		log.Println("Health check")
		nc.Publish(m.Reply, []byte("OK"))
	})

	nc.Subscribe("register", func(m *nats.Msg) {
		var user database.User
		err := json.Unmarshal(m.Data, &user)
		log.Println("Mensaje recibido, usuario: ", user)
		if err != nil {
			log.Println("Error al decodificar el mensaje de registro:", err)
			nc.Publish(m.Reply, []byte("Error en los datos de registro"))
			return
		}

		// Crear el usuario en la base de datos
		userID, err := database.CreateUser(db, user)
		if err != nil {
			log.Println("Error al crear el usuario:", err)
			nc.Publish(m.Reply, []byte("Error al crear el usuario"))
			return
		}

		log.Printf("Usuario creado con ID: %d\n", userID)
		nc.Publish(m.Reply, []byte("Usuario creado con éxito"))
	})

	nc.Subscribe("login", func(m *nats.Msg) {
		// Lógica de inicio de sesión
		var user database.User
		err := json.Unmarshal(m.Data, &user)
		if err != nil {
			log.Println("Error al decodificar el mensaje de inicio de sesión:", err)
			nc.Publish(m.Reply, []byte("Error en los datos de inicio de sesión"))
			return
		}

		usercheck, err := database.LoginUser(db, user)
		if err != nil {
			log.Println("Error al iniciar sesión:", err)
			nc.Publish(m.Reply, []byte("Error al iniciar sesión"))
			return
		}

		token, err := createToken(usercheck)
		if err != nil {
			log.Println("Error al crear el token:", err)
			nc.Publish(m.Reply, []byte("Error al iniciar sesión"))
			return
		}

		nc.Publish(m.Reply, []byte(token))

	})

	nc.Subscribe("check", func(m *nats.Msg) {
		token, err := jwt.Parse(string(m.Data), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			nc.Publish(m.Reply, []byte("invalid"))
			return
		}

		nc.Publish(m.Reply, []byte("valid"))
	})

	log.Println("Authentication service is running and listening on NATS")
	select {}
}
