package models

//ViewedSubject type
type ViewedSubject struct {
	ID        int     `json:"id"`
	IDUsuario int     `json:"id_usuario"`
	Semestre  string  `json:"semestre"`
	Nombre    string  `json:"nombre"`
	Nota      float64 `json:"nota"`
}

// ViewedSubjects array struct
type ViewedSubjects []ViewedSubject
