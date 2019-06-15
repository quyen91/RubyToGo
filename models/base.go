package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

var db *gorm.DB //database

const (
	db_name = "gocontacts"
	db_user = "acvq"
	db_host = "localhost"
	port    = "8000"
)

func DB() *gorm.DB {
 return db
}

// Connect database
func init() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", db_host, db_user, db_name) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{})
}