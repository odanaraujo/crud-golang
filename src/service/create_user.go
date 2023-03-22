package service

import (
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/configuration/rest_err"
	"github.com/odanaraujo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (user *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init createUser model", zap.String("journey ", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepository, err := user.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("error trying to call repository",
			err,
			zap.String("journey", "CreateUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "CreateUser"))
	return userDomain, nil
}
