package token

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type AuthCheck struct {
	UserId   int
	Username string
	Email    string
}

var redisClient *redis.Client
var jwtSecret string

func Init(rc *redis.Client) {
	redisClient = rc
	jwtSecret = os.Getenv("JWT_SECRET")
}

func CreateToken(auth AuthCheck) (string, error) {
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
