package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

	product1 := models.Product{
		ID:               "0001",
		Name:             "Gopro",
		Description:      "This is the description for gopro",
		Images:           []string{"something", "something else"},
		Value:            26000,
		Weight:           0.4,
		UnavailableDates: []time.Time{},
		TotalRating:      3.8,
	}
	product2 := models.Product{
		ID:               "0002",
		Name:             "Gimbal",
		Description:      "This is the description for Gimbal",
		Images:           []string{},
		Value:            32000,
		Weight:           0.8,
		UnavailableDates: []time.Time{},
		TotalRating:      3.8,
	}
	product3 := models.Product{
		ID:               "0003",
		Name:             "Helmet",
		Description:      "This is the description for Helmet",
		Images:           []string{},
		Value:            11000,
		Weight:           1.5,
		UnavailableDates: []time.Time{},
		TotalRating:      4.8,
	}
	product4 := models.Product{
		ID:               "0004",
		Name:             "Trekking Bag",
		Description:      "This is the description for Trekking Bag",
		Images:           []string{},
		Value:            1800,
		Weight:           0.7,
		UnavailableDates: []time.Time{},
		TotalRating:      3.2,
	}
	product5 := models.Product{
		ID:               "0005",
		Name:             "Trekking Pole",
		Description:      "This is the description for Trekking Pole",
		Images:           []string{},
		Value:            1200,
		Weight:           0.6,
		UnavailableDates: []time.Time{},
		TotalRating:      4.2,
	}

	db.Create(&product1)
	db.Create(&product2)
	db.Create(&product3)
	db.Create(&product4)
	db.Create(&product5)

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetProduct)

	router.Run(":" + port)
}
