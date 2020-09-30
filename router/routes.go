package router

import (
	"fmt"
	"net/http"

	"github.com/csepulvedaz/FroidPromedios/service"
	"github.com/gorilla/mux"
)

func indexRouteController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API!!")
}

//Routes creates a router
func Routes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	//Routes
	r.HandleFunc("/", indexRouteController)
	r.HandleFunc("/subject", service.GetSubjects).Methods("GET")
	r.HandleFunc("/subject/viewed", service.CreateViewedSubjects).Methods("POST")

	return r
}
