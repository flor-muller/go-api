package store

import (
	"muller-odontologia/internal/domain"
)

type StoreInterface interface {

	//--------CRUD ODONTOLOGOS--------

	// CreateOdontologo agrega un nuevo odontologo
	CreateOdontologo(odontologo domain.Odontologo) error
	// ReadOdontologo devuelve un odontologo por su id
	ReadOdontologo(id int) (domain.Odontologo, error)
	// UpdateOdontologo actualiza un odontologo
	UpdateOdontologo(odontologo domain.Odontologo) error
	// DeleteOdontologo elimina un odontologo
	DeleteOdontologo(id int) error
	// ExistsMatricula verifica si la matricula de odontologo ya existe
	ExistsMatricula(matricula string) bool

	//--------CRUD PACIENTES--------

	// CreatePaciente agrega un nuevo paciente
	CreatePaciente(paciente domain.Paciente) error
	// ReadPaciente devuelve un paciente por su id
	ReadPaciente(id int) (domain.Paciente, error)
	// UpdatePaciente actualiza un paciente
	UpdatePaciente(paciente domain.Paciente) error
	// DeletePaciente elimina un paciente
	DeletePaciente(id int) error
	// ExistsDni verifica si un DNI de paciente ya existe
	ExistsDni(dni string) bool

	//--------CRUD TURNOS--------

	// CreateTurno agrega un nuevo turno
	CreateTurno(turno domain.Turno) error
	// ReadTurno devuelve un turno por su id
	ReadTurno(id int) (domain.Turno, error)
	// Update actualiza un turno
	UpdateTurno(turno domain.Turno) error
	// Delete elimina un turno
	DeleteTurno(id int) error
	//CreateTurnoDniMatricula agrega un nuevo turno por DNI del paciente y matrícula del dentista
	CreateTurnoDniMatricula(turnoDM domain.TurnoDM) error
	//ReadTurnoDni devuelve turno por DNI del paciente
	ReadTurnoDni(dni string) ([]domain.TurnoDetalle, error)
}
