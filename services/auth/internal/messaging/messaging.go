package messaging

import (
	"auth/internal/logger"
	"auth/internal/token"
	"auth/shared"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var dbUrl string = os.Getenv("DB_SERVICE_URL")

func init() {
	log = logger.GetLogger()
}

func SetupHTTPServer() {
	gin.ForceConsoleColor()
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.POST("/register", handleRegister)
	router.POST("/login", handleLogin)
	router.GET("/health", handleHealthCheck)

	port := os.Getenv("HTTP_PORT")

	http.ListenAndServe(":"+port, router)
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleRegister(c *gin.Context) {
	var user shared.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}
	resp, err := http.Post(dbUrl+"/users/create", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Error("Error on user creation request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on user creation request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleLogin(c *gin.Context) {
}

func SetupSubscribers(nc *nats.Conn) {

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
