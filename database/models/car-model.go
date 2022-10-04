package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Model    string             `json:"model,omitempty" bson:"model,omitempty" validate:"required"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty" validate:"required"`
	Color    string             `json:"color,omitempty" bson:"color,omitempty" validate:"required"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty" validate:"required"`
	BuyValue float64            `json:"buy_value,omitempty" bson:"buy_value,omitempty" validate:"required"`
	DoorsQty int                `json:"doors_qty,omitempty" bson:"doors_qty,omitempty" validate:"required,gt=1,lt=6"`
	SeatsQty int                `json:"seats_qty,omitempty" bson:"seats_qty,omitempty" validate:"required,gt=1,lt=9"`
}

func (c *Car) Validate() error {
	vld := validator.New()
	err := vld.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
