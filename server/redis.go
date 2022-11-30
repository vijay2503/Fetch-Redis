package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var C context.Context
var Client *redis.Client

func Redis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	C = Client.Context()
}

func (LocalService) SetRedis(setkey, value string) (string, error) {
	result, err := Client.Set(C, setkey, value, 0).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func (LocalService) GetRedis(getKey string) (string, error) {
	v, err := Client.Get(C, getKey).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}
func (LocalService) DelRedis(delKey string) error {
	t, err := Client.Del(C, delKey).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if t == 0 {
		return errors.New("Key")
	}
	return nil
}
