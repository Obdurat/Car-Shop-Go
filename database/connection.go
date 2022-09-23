package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)
var client *mongo.Client
var err error

func Connect(name string) *mongo.Collection {
	cliOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(Ctx, cliOpts)
	if err != nil {
		fmt.Println(err.Error())
	}
	DB := client.Database("CarShop").Collection(name)
	return DB
}
