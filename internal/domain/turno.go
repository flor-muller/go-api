package domain

type Turno struct {
	Id           int    `json:"id"`
	IdPaciente   int    `json:"id_paciente"`
	IdOdontologo int    `json:"id_odontologo"`
	Fecha        string `json:"fecha" binding:"required"`
	Hora         string `json:"hora" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
}

// Estructura para ingresar informacion en metodo POST CreateTurnoDniMatricula
type TurnoDM struct {
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Matricula   string `json:"matricula" binding:"required"`
	Dni         string `json:"dni" binding:"required"`
}

// Estructura para devolver informacion en metodo GET ReadTurnoDni
type TurnoDetalle struct {
	Fecha       string     `json:"fecha"`
	Hora        string     `json:"hora"`
	Descripcion string     `json:"descripcion"`
	Paciente    Paciente   `json:"paciente"`
	Odontologo  Odontologo `json:"odontologo"`
}
