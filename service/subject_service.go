package service

import (
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

var subject models.Subject
var subjects models.Subjects

// GetSubjects data
func GetSubjects(w http.ResponseWriter, r *http.Request) {

	subjects, err := repository.GetSubjects()

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, subjects)
}
