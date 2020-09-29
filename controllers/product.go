package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../models"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// taskID, err := strconv.Atoi(vars["id"])
}

//PostProductController route '/' controller with POST method
func PostProductController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := models.GetDB().Prepare("INSERT INTO supermarket_ms_product(name, description, unit_measurement, quantity, category_id) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	description := keyVal["description"]
	unitMeasurement := keyVal["unit_measurement"]
	quantity, _ := strconv.Atoi(keyVal["quantity"])
	categoryID, _ := strconv.Atoi(keyVal["category_id"])

	_, err = stmt.Exec(name, description, unitMeasurement, quantity, categoryID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("map:", keyVal)
	fmt.Fprintf(w, "New product was created")
}
