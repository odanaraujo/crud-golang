package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/configuration/validation"
	"github.com/odanaraujo/crud-golang/src/model/request"
	"log"
)

func CreateUser(c *gin.Context) {

	userRequest := request.UserRequest{}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("error trying to marshal object, error=%s", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}
