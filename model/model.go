package model

import (
	"github.com/jinzhu/gorm"
)

type ServerData struct {
	gorm.Model
	Key   string `json:"key"`
	Value string `json:"value"`
}
