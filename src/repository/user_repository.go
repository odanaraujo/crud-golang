package repository

import (
	"github.com/odanaraujo/crud-golang/src/configuration/rest_err"
	"github.com/odanaraujo/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
