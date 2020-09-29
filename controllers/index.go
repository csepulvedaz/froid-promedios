package controllers

import (
	"fmt"
	"net/http"
)

//IndexRouteController route '/' controller
func IndexRouteController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API!!")
}
