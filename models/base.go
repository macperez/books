package models

import (
	"fmt"
	"os"

	// Uses some dialects to deal with mariadb
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, dbHost, dbName)

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	db.LogMode(true)
}

// GetDB returns the connection
func GetDB() *gorm.DB {
	return db
}
