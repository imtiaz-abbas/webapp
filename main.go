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

	println("== DONE ==")

	db.CreateTable(&models.User{})
	fmt.Println("CREATED USER TABEL")
	db.CreateTable(&models.Blog{})
	fmt.Println("CREATED BLOG TABEL")
	db.CreateTable(&models.Product{})
	fmt.Println("CREATED PRODUCT TABEL")
	db.CreateTable(&models.Category{})
	fmt.Println("CREATED CATEGORY TABEL")
	db.CreateTable(&models.Event{})
	fmt.Println("CREATED EVENT TABEL")
	db.CreateTable(&models.Order{})
	fmt.Println("CREATED ORDER TABEL")
	db.CreateTable(&models.Review{})
	fmt.Println("CREATED REVIEW TABEL")

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProduct)

	router.Run(":" + port)
}
