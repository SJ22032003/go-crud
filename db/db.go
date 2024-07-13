package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectToDb() *mongo.Client {
	URI := os.Getenv("MONGO_URI")

	conn := options.Client().ApplyURI(URI)

	conn.SetAuth(options.Credential{
		Username: "",
		Password: "",
	})

	client, err := mongo.Connect(context.Background(), conn)

	if err != nil {
		log.Fatal(err)
		panic("CANNOT CONNECT TO DB")
	}

	fmt.Println("successfully connected to DB")
	db = client.Database("go-crud")
	return client

}

func GetCollection(coll string) *mongo.Collection {
	return db.Collection(coll);
}