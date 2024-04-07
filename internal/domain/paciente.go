package domain

type Paciente struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Domicilio string `json:"domicilio"`
	Dni       string `json:"dni" binding:"required"`
	Alta      string `json:"alta"`
}
