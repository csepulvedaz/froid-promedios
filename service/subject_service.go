package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/csepulvedaz/FroidPromedios/models"
	"github.com/csepulvedaz/FroidPromedios/repository"
	"github.com/csepulvedaz/FroidPromedios/utils"
	"github.com/gorilla/mux"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

var subjectPost models.SubjectPost

// GetSubjectByID data
func GetSubjectByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	varID, err := strconv.Atoi(params["id"])

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	subject, err := repository.GetSubjectByID(varID)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, subject)
}

// GetSubjects data
func GetSubjects(w http.ResponseWriter, r *http.Request) {

	subjects, err := repository.GetSubjects()

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, subjects)
}

// UpdateSubject data
func UpdateSubject(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Unmarshal([]byte(body), &subjectPost)
	user, err := repository.UpdateSubject(subjectPost)

	if err != nil {
		utils.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusCreated, user)
}
