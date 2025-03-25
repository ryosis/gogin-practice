package controller

import (
	"gogin-practice/middleware"
	"gogin-practice/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetData(incomingRoutes *gin.Engine, db *gorm.DB) {
	api := incomingRoutes.Group("/get-data", middleware.SecurityConfiguration())
	{
		api.GET("/users", service.GetUsers(db))
		api.GET("/user/:id", service.GetDetailUser(db))
		api.GET("/search", service.SearchDetailUser(db))
	}
}

func SetData(incomingRoutes *gin.Engine, db *gorm.DB) {
	api := incomingRoutes.Group("/data", middleware.SecurityConfiguration())
	{
		api.PUT("/users", service.SetUser(db))
		api.POST("/loginform", service.LoginForm())
	}
}
