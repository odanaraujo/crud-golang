package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/configuration/logger"
	"github.com/odanaraujo/crud-golang/src/configuration/validation"
	"github.com/odanaraujo/crud-golang/src/controller/model/request"
	"github.com/odanaraujo/crud-golang/src/model"
	"github.com/odanaraujo/crud-golang/src/view"
	"go.uber.org/zap"
	"net/http"
)

// CreateUser
// CreateUser godoc
// @Summary Create user
// @Tags User
// @Text Create user
// @Param body body request.UserRequest true "Create user request body"
// @Accept json
// @Produce json
// @Success 200 {object} response.UserResponse
// @Router /user [post]
func (userController *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("init create user controller", zap.String("journey", "CreateUser"))

	userRequest := request.UserRequest{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("about to start application", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	userDomain := model.NewUserDomain(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.Age)
	domainResult, err := userController.service.CreateUser(userDomain)

	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
