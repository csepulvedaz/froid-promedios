package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/csepulvedaz/FroidPromedios/models"
	"github.com/csepulvedaz/FroidPromedios/repository"
	"github.com/csepulvedaz/FroidPromedios/utils"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

var viewedSubjectPost models.ViewedSubjectPost
var viewedSubjects models.ViewedSubjects

//CreateViewedSubjects ...
func CreateViewedSubjects(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(body), &viewedSubjectPost)

	user, err := repository.CreateViewedSubjects(viewedSubjectPost)
	if err != nil {
		utils.ResponseWithError(w, http.StatusConflict, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusCreated, user)
}
