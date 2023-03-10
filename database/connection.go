package database

import (
	"fmt"
	_ "gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
//Opening db connection
func ConnectDB()*gorm.DB {
	db, err := gorm.Open("mysql", "root:Karanidancan120@gmail.com@tcp(localhost:3306)/db2?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("Connected...")
	return db
}
