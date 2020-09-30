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

// GetViewedSubjectByID ...
func GetViewedSubjectByID(id int) (models.ViewedSubject, error) {
	db := config.GetDB()

	var viewedSubject models.ViewedSubject

	result, err := db.Query("SELECT * FROM vista_materias_vistas WHERE id = ?", id)
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return viewedSubject, err
	}

	for result.Next() {
		err := result.Scan(&viewedSubject.ID, &viewedSubject.IDUsuario, &viewedSubject.Semestre, &viewedSubject.Nombre, &viewedSubject.Nota)
		if err != nil {
			return viewedSubject, err
		}
	}

	return viewedSubject, nil
}

//CreateViewedSubjects ...
func CreateViewedSubjects(vst models.ViewedSubjectPost) (models.ViewedSubject, error) {
	db := config.GetDB()

	var viewedSubjectPost models.ViewedSubjectPost
	var viewedSubject models.ViewedSubject

	crt, err := db.Prepare("INSERT INTO materia_vista (id_semestre, id_usuario, id_materia, nota) values (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	res, err := crt.Exec(vst.IDSemestre, vst.IDUsuario, vst.IDMateria, vst.Nota)
	if err != nil {
		panic(err.Error())
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	viewedSubjectPost.ID = int(rowID)

	resval, err := GetViewedSubjectByID(viewedSubjectPost.ID)
	if err != nil {
		return viewedSubject, err
	}

	return resval, nil
}
