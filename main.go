package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csepulvedaz/FroidPromedios/controllers"
	"github.com/csepulvedaz/FroidPromedios/models"
)

func main() {
	models.Setup()
	defer models.CloseDB()
	fmt.Println("Port 3000 is listening")
	log.Fatal(http.ListenAndServe(":3000", controllers.Routes()))

}
