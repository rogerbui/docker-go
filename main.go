package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	r := gin.Default()
	dsn := fmt.Printf("%v:secret@tcp(mysql:3306)/docker-go?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	r.GET("/products", func (c *gin.Context)  {
		var products []Product
		db.Find(&products)
		c.JSON(200, products)
	})

	r.POST("/products", func(c *gin.Context){
		code := c.PostForm("code")
		price, _ := strconv.ParseUint(c.PostForm("price"),  10, 32)
		db.Create(&Product{Code: code, Price: uint(price)})
	})

	r.Run()
}