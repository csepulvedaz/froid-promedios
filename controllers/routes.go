package controllers

import (
	"github.com/gorilla/mux"
)

//Routes creates a router
func Routes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	//Routes
	r.HandleFunc("/", IndexRouteController)
	r.HandleFunc("/product", PostProductController).Methods("POST")

	return r
}
