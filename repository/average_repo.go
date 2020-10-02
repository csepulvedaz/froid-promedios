package repository

import (
	"database/sql"
	"log"

	"github.com/csepulvedaz/FroidPromedios/config"
	"github.com/csepulvedaz/FroidPromedios/utils"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// PA ...
func PA() (float64, error) {
	db := config.GetDB()

	var pa float64
	var aux sql.NullFloat64

	result, err := db.Query("SELECT AVG(nota) FROM froid_historia_academica_db.vista_materias WHERE nota >= 3.0")
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return pa, err
	}

	for result.Next() {
		err := result.Scan(&aux)
		pa = aux.Float64
		if err != nil {
			return pa, err
		}
	}

	round := utils.GradeRounder(pa)

	return round, nil
}

// PAPA ...
func PAPA() (float64, error) {
	db := config.GetDB()

	var papa float64
	var aux sql.NullFloat64

	result, err := db.Query("SELECT SUM(creditos*nota)/SUM(creditos) FROM froid_historia_academica_db.vista_materias WHERE nota >= 0.0")
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return papa, err
	}

	for result.Next() {
		err := result.Scan(&aux)
		papa = aux.Float64
		if err != nil {
			return papa, err
		}
	}
	round := utils.GradeRounder(papa)

	return round, nil
}

//PAPPI ...
func PAPPI(sem string) (float64, error) {
	db := config.GetDB()

	var pappi float64
	var aux sql.NullFloat64

	result, err := db.Query("SELECT SUM(creditos*nota)/SUM(creditos) FROM froid_historia_academica_db.vista_materias WHERE nota >= 0.0 AND semestre =?", sem)
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return pappi, err
	}

	for result.Next() {
		err := result.Scan(&aux)
		pappi = aux.Float64
		if err != nil {
			return pappi, err
		}
	}
	round := utils.GradeRounder(pappi)

	return round, nil
}
