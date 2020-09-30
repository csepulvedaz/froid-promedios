package models

//ViewedSubjectPost type
type ViewedSubjectPost struct {
	ID         int     `json:"id"`
	IDSemestre int     `json:"id_semestre"`
	IDUsuario  int     `json:"id_usuario"`
	IDMateria  int     `json:"id_materia"`
	Nota       float64 `json:"nota"`
}

// ViewedSubjectsPost array struct
type ViewedSubjectsPost []ViewedSubjectPost
