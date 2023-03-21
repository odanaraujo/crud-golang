package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/configuration/validation"
	"github.com/odanaraujo/crud-golang/src/controller/model/request"
	"github.com/odanaraujo/crud-golang/src/model"
	"github.com/odanaraujo/crud-golang/src/model/service"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("init create user controller", zap.String("journey", "CreateUser"))

	userRequest := request.UserRequest{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("about to start application", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	userDomain := model.NewUserDomain(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.Age)
	userService := service.NewUserDomainService()

	if err := userService.CreateUser(userDomain); err != nil {
		c.JSON(err.Code, err)
	}

	c.String(http.StatusOK, "")
}
