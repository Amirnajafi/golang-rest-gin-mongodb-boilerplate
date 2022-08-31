package services

import (
	"golangGinMongoDb/configs"
	"golangGinMongoDb/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todosCollections *mongo.Collection = configs.GetCollection(configs.DB, "todos")

func GetTodos(c *gin.Context) {
	var todos []models.Todos
	results, err := todosCollections.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer results.Close(c)
	for results.Next(c) {
		var singleUser models.Todos
		if err = results.Decode(&singleUser); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		todos = append(todos, singleUser)
	}
	c.JSON(http.StatusOK, todos)
}
func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, "post added")
}
func CreateTodo(c *gin.Context) {
	var todo models.Todos
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	newTodo := models.Todos{
		Id:    primitive.NewObjectID(),
		Title: todo.Title,
	}
	result, err := todosCollections.InsertOne(c, newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todos
	todoId := c.Param("id")
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	objId, _ := primitive.ObjectIDFromHex(todoId)
	update := bson.M{"title": todo.Title}
	result, err := todosCollections.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var updatedUser models.Todos
	if result.MatchedCount == 1 {
		err := todosCollections.FindOne(c, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, updatedUser)
}
func DeleteTodo(c *gin.Context) {
	todoId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(todoId)
	result, err := todosCollections.DeleteOne(c, bson.M{"id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No todo found with the given id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
