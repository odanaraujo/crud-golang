package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/odanaraujo/crud-golang/docs"
	"github.com/odanaraujo/crud-golang/src/routes"
	"log"
)

// @title 	CRUD Service API
// @version	1.0
// @description A CRUD service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /user
func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
