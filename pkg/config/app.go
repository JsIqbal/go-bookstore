// package config

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// var (
// 	db *gorm.DB
// )

// func Connect() {
// 	d, err := gorm.Open("mysql", "jsiqbal-X455LJ:asdf1234/simplerest?charset=utf8&parseTime=True&loc=Local")

// 	if err != nil {
// 		panic(err)
// 	}
// 	db = d
// }

// func GetDB() *gorm.DB {
// 	return db
// }

package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	username := "root"
	password := "your_password"
	host := "localhost"
	port := "3306"
	database := "bookstore"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, database)

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
