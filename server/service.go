package server

import "qwik/model"

type LocalServer interface {
	Insert(data model.ServerData) error
	GET(key string) (error, string)
	Update(key, value string) error
	Delete(delId int) error
	SetRedis(setkey, value string) (string, error)
	GetRedis(getKey string) (string, error)
	DelRedis(delKey string) error
}
