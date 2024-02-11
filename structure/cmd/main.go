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
	
	// // connect to database
	// 	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// 	db := options.Client().ApplyURI("mongodb+srv://Youssef:<password>@cluster0.wotxocw.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// 	// Create a new client and connect to the server
	// 	client, err := mongo.Connect(context.TODO(), db)
	// 	if err != nil {
	// 	  panic(err)
	// 	}
	// 	defer func() {
	// 	  if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	  }
	// 	}()
	// 	// Send a ping to confirm a successful connection
	// 	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	// 	  panic(err)
	// 	}
	// 	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
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
	
	