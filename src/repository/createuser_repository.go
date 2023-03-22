package repository

import (
	"context"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/configuration/rest_err"
	"github.com/odanaraujo/crud-golang/src/entity/converter"
	"github.com/odanaraujo/crud-golang/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func (userRepository *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init create user repository", zap.String("journey", "CreateUser"))

	collectionName := os.Getenv(MONGODB_USER_DB)

	collection := userRepository.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		logger.Error("error trying to create user", err, zap.String("journey", "CreateUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return userDomain, nil
}
