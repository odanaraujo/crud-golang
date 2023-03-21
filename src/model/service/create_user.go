package service

import (
	"fmt"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/configuration/rest_err"
	"github.com/odanaraujo/crud-golang/src/model"
	"go.uber.org/zap"
)

func (user *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init createUser model", zap.String("journey ", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
