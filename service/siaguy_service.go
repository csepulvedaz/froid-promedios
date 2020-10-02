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

// SiaGuy advance
func SiaGuy(w http.ResponseWriter, r *http.Request) {
	idUsuario := r.FormValue("id_usuario")
	id, err := strconv.Atoi(idUsuario)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	siaGuy, err := repository.GetSiaGuy(id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, siaGuy)
}
