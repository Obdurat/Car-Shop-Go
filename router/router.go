package router

import (
	"github.com/Obdurat/Carshop-Go/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/car", controllers.GetAllCars).Methods("GET")
	r.HandleFunc("/car", controllers.CreateCar).Methods("POST")
	r.HandleFunc("/car/{id}", controllers.GetOneCar).Methods("GET")
	r.HandleFunc("/car/{id}", controllers.UpdateCar).Methods("PUT")
	r.HandleFunc("/car/{id}", controllers.DeleteCar).Methods("DELETE")
	return r
}
