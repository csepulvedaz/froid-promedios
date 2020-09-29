package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csepulvedaz/FroidPromedios/controllers"
)

func main() {
	controllers.Setup()
	defer controllers.CloseDB()
	fmt.Println("Port 8080 is listening")
	log.Fatal(http.ListenAndServe(":8080", controllers.Routes()))

}
