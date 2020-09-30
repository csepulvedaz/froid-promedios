package models

//Subject type
type Subject struct {
	ID            int    `json:"id"`
	IDUsuario     int    `json:"id_usuario"`
	Semestre      string `json:"semestre"`
	Tipologia     string `json:"tipologia"`
	Nombre        string `json:"nombre"`
	Prerequisitos string `json:"prerequisitos"`
	Creditos      int    `json:"creditos"`
}

// Subjects array struct
type Subjects []Subject
