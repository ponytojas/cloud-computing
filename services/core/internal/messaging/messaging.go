package messaging

import (
	"bytes"
	"core/internal/logger"
	"core/shared"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var authUrl string

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	log = logger.GetLogger()
	authUrl = os.Getenv("DB_SERVICE_URL")
}

func SetupHTTPServer() {
	gin.ForceConsoleColor()
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/health", handleHealthCheck)
	router.POST("/login", handleLogin)
	router.POST("/logout", handleLogout)
	router.POST("/register", handleRegister)

	port := os.Getenv("HTTP_PORT")

	log.Infof("Core service started on port %s", os.Getenv("HTTP_PORT"))
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}

func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleLogin(c *gin.Context) {
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

	resp, err := http.Post(authUrl+"/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error("Error sending request:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing response body:", err)
		}
	}(resp.Body)

	var token shared.Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		log.Error("Error decoding response body:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token.Token})
}

func handleLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func handleRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
