package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/user/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
