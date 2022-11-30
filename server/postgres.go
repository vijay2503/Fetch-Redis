package server

import (
	"errors"
	"fmt"
	"log"
	"os"
	"qwik/model"
	_ "sync"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var Db *gorm.DB

type LocalService struct{}

func ConnectDrive() {
	var err error
	if err := godotenv.Load(`D:\Fetch-Redis\server\dbconfig.env`); err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	Db, err = Connect(host, port, user, password, dbname)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
func TableCreation() {
	table := model.ServerData{}
	Db.CreateTable(&table).SingularTable(true)
}
func (LocalService) Insert(data model.ServerData) error {
	if err := Db.Table("server_data").Create(&data).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (LocalService) GET(key string) (error, string) {
	serverData := &model.ServerData{}
	err := Db.Table("server_data").Find(serverData, "key=?", key).Error
	if err != nil {
		return err, ""
	}
	return nil, serverData.Value
}
func (LocalService) Delete(delId int) error {
	serverData := &model.ServerData{}
	db := Db.Table("server_data").Where("id=$1", delId).Delete(serverData, "key=?", delId)
	if db.Error != nil {
		log.Println("func Name : RepoDeleteStudentDetails error=", db.Error)
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.New("not exist")
	}
	return nil
}

func (LocalService) Update(key, value string) error {
	serverData := &model.ServerData{}
	err := Db.Table("server_data").Find(serverData, "key=?", key).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	serverData.Key = key
	serverData.Value = value
	if err := Db.Save(serverData).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
