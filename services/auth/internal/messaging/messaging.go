package messaging

import (
	"auth/internal/logger"
	"auth/internal/token"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	log = logger.GetLogger()
}

func SetupHTTPServer() {
	router := gin.Default()

	// Definiendo endpoints
	router.POST("/register", handleRegister)
	router.POST("/login", handleLogin)
	router.GET("/health", handleHealthCheck)

	// Iniciar el servidor en un puerto específico
	router.Run(":8080")
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleRegister(c *gin.Context) {
}

func SetupSubscribers(nc *nats.Conn) {
	nc.Subscribe("register", func(m *nats.Msg) {
		response, err := nc.Request("database.users.create", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error creating user:", err)
			nc.Publish(m.Reply, []byte("Error creating user"))
			return
		}

		if string(response.Data) == "OK" {
			nc.Publish(m.Reply, []byte("OK"))
		} else {
			log.Error("ERROR:", string(response.Data))
			nc.Publish(m.Reply, []byte("Error creating user"))
		}
	})

	nc.Subscribe("login", func(m *nats.Msg) {
		response, err := nc.Request("database.users.login", m.Data, 1000*time.Millisecond)
		if err != nil {
			log.Error("Error al iniciar sesión:", err)
			nc.Publish(m.Reply, []byte("Error at login"))
			return
		}

		if string(response.Data) == "ERROR" {
			nc.Publish(m.Reply, []byte("Error at login"))
			return
		} else {
			log.Debug("Login correcto")
		}

		var usercheck token.AuthCheck
		err = json.Unmarshal(response.Data, &usercheck)
		if err != nil {
			log.Error("Error at login:", err)
			nc.Publish(m.Reply, []byte("Error at login"))
			return
		}

		token, err := token.CreateToken(usercheck)
		if err != nil {
			log.Error("Error al crear el token:", err)
			nc.Publish(m.Reply, []byte("Error at login"))
			return
		}

		nc.Publish(m.Reply, []byte(token))

	})

}
