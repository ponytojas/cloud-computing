package cart

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func Init(rc *redis.Client) {
	redisClient = rc
}

func GetUserCart(userId string) (map[string]string, error) {
	ctx := context.Background()
	cart, err := redisClient.HGetAll(ctx, "cart_"+userId).Result()
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func ClearUserCart(userId string) error {
	ctx := context.Background()
	_, err := redisClient.Del(ctx, "cart_"+userId).Result()
	if err != nil {
		return err
	}
	return nil
}
