package server

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func Connect(host, port, user, password, dbname string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
