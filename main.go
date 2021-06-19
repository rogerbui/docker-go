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

	fmt.Println("Environment: ", os.Getenv("ENVIRONMENT"))
	
	r := gin.Default()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", 
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D421111", Price: 100})

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