package server

import (
	"context"
	"errors"
	"fmt"
	"time"

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

//SetRedis to set the key-value data to the redis sever
func (LocalService) SetRedis(setkey, value string) (string, error) {
	result, err := Client.Set(C, setkey, value, 1*time.Hour).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

//GetRedis Get the key-value pair from redis
func (LocalService) GetRedis(getKey string) (string, error) {
	v, err := Client.Get(C, getKey).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

//DelRedis Delete the single data from redis
func (LocalService) DelRedis(delKey string) error {
	t, err := Client.Del(C, delKey).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if t == 0 {
		return errors.New("Key Does Not Exist")
	}
	return nil
}
