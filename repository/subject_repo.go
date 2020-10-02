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

// GetSubjectByID ...
func GetSubjectByID(id int) (models.Subject, error) {
	db := config.GetDB()

	var subject models.Subject

	result, err := db.Query("SELECT * FROM vista_materias WHERE id = ?", id)
	if err != nil {
		log.Println("Error query user: " + err.Error())
		return subject, err
	}

	for result.Next() {
		err := result.Scan(&subject.ID, &subject.IDUsuario, &subject.Semestre, &subject.Tipologia, &subject.Nombre, &subject.Prerequisitos, &subject.Creditos, &subject.Nota)
		if err != nil {
			return subject, err
		}
	}

	return subject, nil
}

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
		err := rows.Scan(&subject.ID, &subject.IDUsuario, &subject.Semestre, &subject.Tipologia, &subject.Nombre, &subject.Prerequisitos, &subject.Creditos, &subject.Nota)
		if err != nil {
			panic(err.Error())
		} else {
			subjects = append(subjects, subject)
		}
	}

	return subjects, nil
}

//CreateRepeatSubject ...
func CreateRepeatSubject(idUsuario int, idSemestre int, nombre string, creditos int, nota float64) (models.Subject, error) {
	db := config.GetDB()

	var subject models.Subject

	crt, err := db.Prepare(`INSERT INTO materia (id_usuario, id_semestre, id_tipologia, nombre, prerequisitos, creditos, nota) values (?, ?, 4, ?, "", ?, ?)`)
	if err != nil {
		return subject, err
	}
	res, err := crt.Exec(idUsuario, idSemestre, nombre, creditos, nota)
	if err != nil {
		return subject, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		return subject, err
	}

	subject.ID = int(rowID)

	resval, err := GetSubjectByID(subject.ID)
	if err != nil {
		return subject, err
	}

	return resval, nil
}

//UpdateSubject ...
func UpdateSubject(sub models.SubjectPost) (models.Subject, error) {
	db := config.GetDB()

	var result models.Subject

	if sub.Nota > 5.0 {
		sub.Nota = 5.0
	} else if sub.Nota < 0.0 {
		sub.Nota = 0.0
	}

	row, err := GetSubjectByID(sub.ID)
	if err != nil {
		return result, err
	}

	if row.Nota == -1 {
		crt, err := db.Prepare("UPDATE materia SET nota =? where id =?")
		if err != nil {
			panic(err.Error())
		}
		_, queryError := crt.Exec(sub.Nota, sub.ID)
		if queryError != nil {
			panic(err.Error())
		}
		resval, err := GetSubjectByID(sub.ID)
		if err != nil {
			return result, err
		}
		result = resval
	} else {
		resval, err := CreateRepeatSubject(row.IDUsuario, sub.IDSemestre, row.Nombre, row.Creditos, sub.Nota)
		if err != nil {
			return result, err
		}
		result = resval
	}

	return result, nil
}
