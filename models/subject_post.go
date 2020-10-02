package models

//SubjectPost type
type SubjectPost struct {
	ID         int     `json:"id"`
	IDSemestre int     `json:"id_semestre"`
	Nota       float64 `json:"nota"`
}

// SubjectsPost array struct
type SubjectsPost []SubjectPost
