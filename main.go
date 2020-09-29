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
	fmt.Println("Port 3000 is listening")
	log.Fatal(http.ListenAndServe(":3000", controllers.Routes()))

}
