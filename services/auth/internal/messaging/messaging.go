package messaging

import (
	"auth/internal/logger"
	"auth/internal/token"
	"auth/shared"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var dbUrl string

func init() {
	godotenv.Load()
	log = logger.GetLogger()
	dbUrl = os.Getenv("DB_SERVICE_URL")
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

	resp, err := http.Post(dbUrl+"/users/login", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Error("Error on user login request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Error on user login request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var usercheck token.AuthCheck
	err = json.NewDecoder(resp.Body).Decode(&usercheck)
	if err != nil {
		log.Error("Error on user login request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := token.CreateToken(usercheck)
	if err != nil {
		log.Error("Error on user login request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
