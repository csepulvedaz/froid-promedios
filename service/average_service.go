package service

import (
	"net/http"
	"strconv"

	"github.com/csepulvedaz/FroidPromedios/repository"
	"github.com/csepulvedaz/FroidPromedios/utils"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// PA calculate
func PA(w http.ResponseWriter, r *http.Request) {

	idUsuario := r.FormValue("id_usuario")
	id, err := strconv.Atoi(idUsuario)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	pa, err := repository.PA(id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, pa)
}

// PAPA calculate
func PAPA(w http.ResponseWriter, r *http.Request) {
	idUsuario := r.FormValue("id_usuario")
	id, err := strconv.Atoi(idUsuario)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	papa, err := repository.PAPA(id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, papa)
}

// PAPPI calculate
func PAPPI(w http.ResponseWriter, r *http.Request) {
	semestre := r.FormValue("semestre")
	idUsuario := r.FormValue("id_usuario")

	id, err := strconv.Atoi(idUsuario)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	pappi, err := repository.PAPPI(semestre, id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, pappi)
}
