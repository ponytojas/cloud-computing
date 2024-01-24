package sharedcart

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var jwtSecret string

func Init(rc *redis.Client) {
	redisClient = rc
}

func AddToUserCart(userId string, productId string, quantity int64) error {
	ctx := context.Background()
	_, err := redisClient.Exists(ctx, "cart_"+userId).Result()
	if err != nil {
		return err
	}
	if err == redis.Nil {
		_, err := redisClient.HSet(ctx, "cart_"+userId, productId, quantity).Result()
		if err != nil {
			return err
		}
		return nil
	}

	_, err = redisClient.HSet(ctx, "cart_"+userId, productId, quantity).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetUserCart(userId string) (map[string]string, error) {
	ctx := context.Background()
	cart, err := redisClient.HGetAll(ctx, "cart_"+userId).Result()
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func ClearCart(userId string) error {
	ctx := context.Background()
	_, err := redisClient.Del(ctx, "cart_"+userId).Result()
	if err != nil {
		return err
	}
	return nil
}
