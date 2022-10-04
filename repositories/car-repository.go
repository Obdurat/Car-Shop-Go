package repositories

import (
	"context"
	"encoding/json"

	"github.com/Obdurat/Carshop-Go/database"
	"github.com/Obdurat/Carshop-Go/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DB = database.Connect("Vehicles")

func GetAll(ctx context.Context) ([]models.Car, error) {
	var cars []models.Car
	cursor, err := DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var car models.Car
		cursor.Decode(&car)
		cars = append(cars, car)
	}
	return cars, nil
}

func GetOne(id primitive.ObjectID, ctx context.Context) (interface{}, error) {
	var car models.Car
	err := DB.FindOne(ctx, bson.M{"_id": id}).Decode(&car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func Create(ctx context.Context, body []byte) (interface{}, error) {
	var car models.Car
	json.Unmarshal(body, &car)
	err := car.Validate()
	if err != nil { return nil, err }
	r, err := DB.InsertOne(ctx, car)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Update(ctx context.Context, id primitive.ObjectID, body []byte) (interface{}, error) {
	var car models.Car
	json.Unmarshal(body, &car)
	r, err := DB.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": car})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Delete(ctx context.Context, id primitive.ObjectID) (interface{}, error) {
	r, err := DB.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return r, nil
}
