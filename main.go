package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/webapp/controllers"
	"github.com/webapp/models"
)

func dbURL() string {
	var sslmode string
	sslmode = "sslmode=disable"
	dbPort, _ := strconv.Atoi(os.Getenv("DBPORT"))
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s %s",
		os.Getenv("DBHOST"),
		dbPort,
		os.Getenv("DBUSER"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPASSWORD"),
		sslmode,
	)
}

func main() {
	port := os.Getenv("PORT")
	db, err := gorm.Open("postgres", dbURL())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dbase := db.DB()
	defer dbase.Close()
	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}
	if port == "" {
		port = "8080"
	}
	db.DropTableIfExists(&models.User{})
	fmt.Println("DELETED USER TABEL")
	db.DropTableIfExists(&models.Blog{})
	fmt.Println("DELETED BLOG TABEL")
	db.DropTableIfExists(&models.Product{})
	fmt.Println("DELETED PRODUCT TABEL")
	db.DropTableIfExists(&models.Category{})
	fmt.Println("DELETED CATEGORY TABEL")
	db.DropTableIfExists(&models.Event{})
	fmt.Println("DELETED EVENT TABEL")
	db.DropTableIfExists(&models.Order{})
	fmt.Println("DELETED ORDER TABEL")
	db.DropTableIfExists(&models.Review{})
	fmt.Println("DELETED REVIEW TABEL")

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProduct)

	router.Run(":" + port)
}
