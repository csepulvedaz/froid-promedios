package repository

import (
	"database/sql"
	"log"

	"github.com/csepulvedaz/FroidPromedios/config"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// GetSiaGuy ...
func GetSiaGuy(id int) (float64, error) {
	db := config.GetDB()

	var approved, total float64
	var auxA, auxT sql.NullFloat64

	a, err := db.Query("SELECT SUM(creditos) AS creditos_aprobados FROM froid_historia_academica_db.vista_materias WHERE nota >= 3.0")
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return approved, err
	}

	for a.Next() {
		err := a.Scan(&auxA)
		approved = auxA.Float64
		if err != nil {
			return approved, err
		}
	}

	t, err := db.Query(`SELECT SUM(creditos) AS total_creditos FROM froid_historia_academica_db.vista_materias WHERE tipologia != "Fantasma"`)
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return total, err
	}

	for t.Next() {
		err := t.Scan(&auxT)
		total = auxT.Float64
		if err != nil {
			return total, err
		}
	}

	advance := (approved / total) * 100

	return advance, nil
}
