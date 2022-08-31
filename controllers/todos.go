package controllers

import (
	"golangGinMongoDb/services"

	"github.com/gin-gonic/gin"
)

// add TodoController
func AddTodoRoutes(api *gin.RouterGroup) {
	api.GET("/todos", services.GetTodos)
	api.POST("/todos", services.CreateTodo)
	api.GET("/todos/:id", services.GetTodo)
	api.PUT("/todos/:id", services.UpdateTodo)
	api.DELETE("/todos/:id", services.DeleteTodo)
}
