package database
import(
"fmt"
"log"
"os"
"context"
"github.com/joho/godotenv"
"go.mongo.org/mongo-driver/mongo"
"go.mongo.org/mongo-driver/mongo/options"
)
func DBinstance() *mongo.Client{
//connect to database
	err:= godotenv.Load(".env")
	if err != nil {
		log.Fatal("ERROR Loading .env File")
	}
	MongoDb= os.Getenv("MONGODB_URL")
		// // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
		// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		// db := options.Client().ApplyURI("mongodb+srv://Youssef:<password>@cluster0.wotxocw.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
		// // Create a new client and connect to the server
		// client, err := mongo.Connect(context.TODO(), db)
		// if err != nil {
		//   panic(err)
		// }
		// defer func() {
		//   if err = client.Disconnect(context.TODO()); err != nil {
		// 	panic(err)
		//   }
		// }()
		// // Send a ping to confirm a successful connection
		// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		//   panic(err)
		// }
		// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}