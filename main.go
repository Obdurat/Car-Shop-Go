package main

import (
	"fmt"
	"net/http"

	"github.com/Obdurat/Carshop-Go/router"
)

func main() {
	fmt.Println("Starting api ...")
	routes := router.Router()
	http.ListenAndServe(":8080", routes)
}
