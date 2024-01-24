package token

import (
	"bytes"
	"cart/internal/logger"
	"cart/shared"
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
	godotenv.Load()
	log = logger.GetLogger()
	authUrl = os.Getenv("AUTH_SERVICE_URL")
}

func CheckTokenMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	token := shared.Token{Token: authHeader}

	requestBody, err := json.Marshal(token)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return
	}

	resp, err := http.Post(authUrl+"/check", "application/json", bytes.NewBuffer(requestBody))
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

	var response shared.TokenCheckResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error("Error decoding response body:", err)
		return
	}

	if !response.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
