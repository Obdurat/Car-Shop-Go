package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Model    string             `json:"model,omitempty" bson:"model,omitempty"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty"`
	Color    string             `json:"color,omitempty" bson:"color,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
	BuyValue float64            `json:"buy_value,omitempty" bson:"buy_value,omitempty"`
	DoorsQty int                `json:"doors_qty,omitempty" bson:"doors_qty,omitempty"`
	SeatsQty int                `json:"seats_qty,omitempty" bson:"seats_qty,omitempty"`
}

func (c *Car) Validate() error {
	fmt.Println("Validate")
	return nil
}
