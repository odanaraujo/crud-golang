package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/crud-golang/src/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/user/:userEmail", controller.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
