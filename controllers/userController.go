package controllers

import (
	"context"
	"example/structure/database"
	"example/structure/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"organization-management-api/helpers"
	helper "organization-management-api/helpers"
	"organization-management-api/models"

	"golang.org/x/crypto/bcrypt"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.New()
func HashPassword()

func VerifyPassword()

func Signup()  gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(),100*(time.Second))
		var user models.User
		if err := c.BindJSON((&user)); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}
		if validationErr:= validate.Struct(user); validationErr!= nil{
			c.JSON(http.StatusBadRequest, gin.H{"error" : validationErr.Error()})
			return
		}
		count,err :=userCollection.CountDocuments(ctx,bson.M{"email": user.Email})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "couldn't find your email"})
			return
		}
		if count >0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "This email already exists"})
			return
		}
	}
}

func Signin()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")
	if err:=	helper.MatchUserTypetoUid(c, userId); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})

		return
	}
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
	var user  models.User
	err :=userCollection.FindOneAndDelete(ctx,bson.M{"user_id" : userId}).Decode(&user)
	defer cancel()
	if err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK,user)

	}
}