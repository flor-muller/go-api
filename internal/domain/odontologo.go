package domain

type Odontolgo struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}
