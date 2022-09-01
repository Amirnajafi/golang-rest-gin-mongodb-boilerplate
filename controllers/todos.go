package controllers

import (
	"golangGinMongoDb/middlewares"
	"golangGinMongoDb/services"

	"github.com/gin-gonic/gin"
)

// add TodoController
func AddTodoRoutes(api *gin.RouterGroup) {
	api.GET("/todos", middlewares.WithAuthentication(), services.GetTodos)
	api.POST("/todos", middlewares.WithAuthentication(), services.CreateTodo)
	api.GET("/todos/:id", middlewares.WithAuthentication(), services.GetTodo)
	api.PUT("/todos/:id", middlewares.WithAuthentication(), services.UpdateTodo)
	api.DELETE("/todos/:id", middlewares.WithAuthentication(), services.DeleteTodo)
}
