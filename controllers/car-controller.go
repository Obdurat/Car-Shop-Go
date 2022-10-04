package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Obdurat/Carshop-Go/repositories"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.Background()

func CreateCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	defer req.Body.Close()

	b, err := io.ReadAll(req.Body)
	if err != nil {	SendError(err, res, http.StatusBadRequest);	return }

	i, err := repositories.Create(ctx, b)
	if err != nil {	SendError(err, res, http.StatusBadRequest);	return }

	defer json.NewEncoder(res).Encode(i)
}

func GetAllCars(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	cars, err := repositories.GetAll(ctx)
	if err != nil { SendError(err, res, http.StatusBadRequest);	return }

	defer json.NewEncoder(res).Encode(cars)
}

func GetOneCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	params := mux.Vars(req)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {	SendError(err, res, http.StatusBadRequest); return }

	car, err := repositories.GetOne(id, ctx)
	if err != nil {	SendError(err, res, http.StatusBadRequest);	return }

	defer json.NewEncoder(res).Encode(car)
}

func UpdateCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	params := mux.Vars(req)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {	SendError(err, res, 400); return }

	b, err := io.ReadAll(req.Body)
	if err != nil {	SendError(err, res, 400); return }

	car, err := repositories.Update(ctx, id, b)
	if err != nil {	SendError(err, res, 404); return }

	defer json.NewEncoder(res).Encode(car)
}

func DeleteCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {	SendError(err, res, 400); return }

	car, err := repositories.Delete(ctx, id)
	if err != nil {	SendError(err, res, 404); return }

	defer json.NewEncoder(res).Encode(car)
}

func SendError(err error, res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	res.Write([]byte(err.Error()))
}
