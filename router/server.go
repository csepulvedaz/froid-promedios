package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csepulvedaz/FroidPromedios/config"
)

//Server initialize
func Server(port string) {
	config.Connect()
	defer config.CloseDB()
	fmt.Printf("Port %s is listening.\n", port)
	log.Fatal(http.ListenAndServe(":"+port, Routes()))
}
