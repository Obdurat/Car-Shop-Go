package router

import (
	"github.com/Obdurat/Carshop-Go/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetAllCars).Methods("GET")
	r.HandleFunc("/", controllers.CreateCar).Methods("POST")
	r.HandleFunc("/{id}", controllers.GetOneCar).Methods("GET")
	r.HandleFunc("/{id}", controllers.UpdateCar).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteCar).Methods("DELETE")
	return r
}
