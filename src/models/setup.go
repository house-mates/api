package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	dsn := "root:rootpass@tcp(127.0.0.1:3306)/dbname?charset=latin1&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&User{},
		&Usermeta{},
		&Action{},
		&Event{},
		&Response{},
	)

	DB = database
}
