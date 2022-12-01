package server

import (
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

//ConnectDrive connect With Local Data-Base
func ConnectDrive() {
	var err error
	if err := godotenv.Load(`D:\Fetch-Redis\server\dbconfig.env`); err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	Db, err = Connect(os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	if err != nil {
		fmt.Println(err)
	}
}

//Table Creation to crete the table to the database
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

//GET to get the data from local-database
func (LocalService) GET(key string) (error, string) {
	serverData := &model.ServerData{}
	err := Db.Table("server_data").Find(serverData, "key=?", key).Error
	if err != nil {
		return err, ""
	}
	return nil, serverData.Value
}

//Update is Update the data to the local data-base
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
