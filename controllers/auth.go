package controllers

import (
	"golangGinMongoDb/services"

	"github.com/gin-gonic/gin"
)

// add TodoController
func AuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", services.Login)
	api.POST("/register", services.Register)
}
