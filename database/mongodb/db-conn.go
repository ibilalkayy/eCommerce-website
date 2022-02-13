package mongodb

// Importing the libraries
import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connect = Connection()

// It will build the connection with the mongodb database.
func Connection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster1.2inap.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
