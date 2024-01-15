package token

import (
	"auth/shared"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var jwtSecret string

func Init(rc *redis.Client) {
	redisClient = rc
	jwtSecret = os.Getenv("JWT_SECRET")
}

func CreateToken(auth shared.AuthCheck) (string, error) {
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

func CheckToken(c *gin.Context) {
	var tokenString string
	err := c.ShouldBindJSON(&tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(http.StatusOK, gin.H{"claims": claims})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
}

func DeleteToken(tokenString string) error {
	ctx := context.Background()
	_, err := redisClient.Del(ctx, tokenString).Result()
	if err != nil {
		return err
	}

	return nil
}
