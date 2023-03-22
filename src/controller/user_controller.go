package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{serviceInterface}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
