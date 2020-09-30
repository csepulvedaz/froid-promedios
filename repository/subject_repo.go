package repository

import (
	"log"

	"github.com/csepulvedaz/FroidPromedios/config"
	"github.com/csepulvedaz/FroidPromedios/models"

	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	_ "github.com/go-sql-driver/mysql"
)

// GetSubjects ...
func GetSubjects() (models.Subjects, error) {
	db := config.GetDB()

	var subject models.Subject
	var subjects models.Subjects

	rows, err := db.Query("SELECT * FROM vista_materias")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err := rows.Scan(&subject.ID, &subject.IDUsuario, &subject.Semestre, &subject.Tipologia, &subject.Nombre, &subject.Prerequisitos, &subject.Creditos)
		if err != nil {
			panic(err.Error())
		} else {
			subjects = append(subjects, subject)
		}
	}

	return subjects, nil
}
