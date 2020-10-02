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
	//Subject
	r.HandleFunc("/subject", service.GetSubjects).Methods("GET")
	r.HandleFunc("/subject/{id}", service.GetSubjectByID).Methods("GET")
	r.HandleFunc("/subject", service.UpdateSubject).Methods("PUT")
	//Averages
	r.HandleFunc("/average/pa", service.PA).Methods("GET")
	r.HandleFunc("/average/papa", service.PAPA).Methods("GET")
	r.HandleFunc("/average/pappi", service.PAPPI).Methods("GET")
	r.HandleFunc("/sia-guy", service.SiaGuy).Methods("GET")

	return r
}
