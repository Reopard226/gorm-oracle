package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//Simple program to test GORM on Oracle DB
func main() {
	fmt.Println("opening connection to oracle...")
	dsn := "oracle://system:Oradoc_db1@localhost:1521/ORCLCDB.localdomain"
	db, err := gorm.Open("goracle", dsn)
	if err != nil {
		fmt.Println(err)
		log.Fatal("failed to connect database")
	}
	fmt.Println("Connection opened...")
	db.LogMode(true)
	defer db.Close()

	db.DropTable(&Product{})
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "ABC", Price: 10})

	var product Product
	db.First(&product, 1)
	log.Println(product)

}
