package main

import (
	"golangGinMongoDb/configs"
	"golangGinMongoDb/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getRoutes() {
	api := router.Group("/api")
	controllers.AddTodoRoutes(api)
}

func Run() {
	getRoutes()
	configs.ConnectDB()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Run(":5000")
}

func main() {
	Run()
}
