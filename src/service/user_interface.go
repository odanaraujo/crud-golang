package service

import (
	"github.com/odanaraujo/crud-golang/src/configuration/rest_err"
	"github.com/odanaraujo/crud-golang/src/model"
	"github.com/odanaraujo/crud-golang/src/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, domainInterface model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
