package main

import (
	routes "golang-jwt-project/routes"
	"os"

	"github.com/gin-gonic/gin"
)

// type User struct{
// 	ID int		 `json:"id"`
// 	Name 	string	`json:"name"`
// 	Email 	string `json:"email"`
// }
func main(){
	// Setting up a Port
	port := os.Getenv("PORT")
	if port =="" {
		port ="8000"
	}
	
	//create router
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.GET("/api-1",func(c *gin.Context){
		c.JSON(200, gin.H{"success" :"Access granted for api-1"})
	})
	router.GET("/api-2",func(c *gin.Context){
		c.JSON(200, gin.H{"success" :"Access granted for api-2"})
	})
	router.Run(":"+ port)


	
	}
	
	