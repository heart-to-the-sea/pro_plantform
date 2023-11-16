package router

import (
	"server/controllers"
	"server/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	user := r.Group("/user")
	{
		user.GET("/login", controllers.UserController{}.GetLogin)
		user.GET("/list", controllers.UserController{}.List)
		user.GET("/add", controllers.UserController{}.Add)
		user.GET("/delete", controllers.UserController{}.Delete)
		user.GET("/testException", controllers.UserController{}.Exception)
	}
	return r
}
