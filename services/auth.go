package services

import (
	"golangGinMongoDb/configs"
	"golangGinMongoDb/dto"
	"golangGinMongoDb/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollections *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func Login(c *gin.Context) {
	var user dto.RegisterDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}
	var findedUser models.Users
	result := userCollections.FindOne(c, bson.M{"email": user.Email})
	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user or password is incorrect"})
		return
	}
	if err := result.Decode(&findedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(findedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password is incorrect"})
		return
	}
	// create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": findedUser.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": tokenString})

}

func Register(c *gin.Context) {
	var user dto.RegisterDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}

	count, err := userCollections.CountDocuments(c, bson.M{"email": user.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	newUser := models.Users{
		Id:       primitive.NewObjectID(),
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	_, err = userCollections.InsertOne(c, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
