package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongo.org/mongo-driver/mongo"
	"go.mongo.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func DBinstance() *mongo.Client{
//connect to database
	err:= godotenv.Load(".env")
	if err != nil {
		log.Fatal("ERROR Loading .env File")
	}
	// Getting Connection String
	MongoDb := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		  log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(),10 *time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
 	 }		
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
		return client
}
 // Initializing a DB instance
var Client *mongo.Client = DBinstance()
func OpenCollection(client *mongo.Client, CollectionName string) *mongo.Collection{
	var collection *mongo.collection= client.Database("cluster0").Collection(CollectionName)
	return collection

}