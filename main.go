package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
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
	dsn := fmt.Sprintf("sqlserver://%v:%v@%v:1433", 
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"))
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE %v", os.Getenv("DB_NAME")))
	db.Exec(fmt.Sprintf("USE %v", os.Getenv("DB_NAME")))

	// Migrate the schema
	db.AutoMigrate(&Product{})

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