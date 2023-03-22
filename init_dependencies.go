package main

import (
	"github.com/odanaraujo/crud-golang/src/controller"
	"github.com/odanaraujo/crud-golang/src/repository"
	"github.com/odanaraujo/crud-golang/src/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	serviceDomain := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(serviceDomain)
	return userController
}
