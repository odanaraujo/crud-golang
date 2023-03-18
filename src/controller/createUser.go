package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/configuration/validation"
	"github.com/odanaraujo/crud-golang/src/model/request"
	"github.com/odanaraujo/crud-golang/src/model/response"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("init create user controller")

	userRequest := request.UserRequest{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("about to start application", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	userResponse := response.UserResponse{
		ID:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, userResponse)
}
