package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/odanaraujo/crud-golang/docs"
	"github.com/odanaraujo/crud-golang/src/configuration/database/mongodb"
	"github.com/odanaraujo/crud-golang/src/routes"
	"log"
)

// @title 	CRUD Service API
// @version	1.0
// @description A CRUD service API in Go using Gin framework
// @host 	localhost:8080
func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongodbConnection(context.Background())

	if err != nil {
		log.Fatalf("error trying to connect to database error=%s", err.Error())
		return
	}

	userController := initDependencies(database)

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup, userController)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
