package store

import (
	"muller-odontologia/internal/domain"
)

type StoreInterface interface {

	//--------CRUD ODONTOLOGOS--------

	// Create agrega un nuevo odontologo
	CreateOdontologo(odontologo domain.Odontologo) error
	// Read devuelve un odontologo por su id
	ReadOdontologo(id int) (domain.Odontologo, error)
	// Update actualiza un odontologo
	UpdateOdontologo(odontologo domain.Odontologo) error
	// Delete elimina un odontologo
	DeleteOdontologo(id int) error

	// Exists verifica si la matricula de odontologo ya existe
	ExistsMatricula(matricula string) bool

	//GetAll() ([]domain.Odontologo, error)

	//--------CRUD PACIENTES--------

	// Create agrega un nuevo paciente
	CreatePaciente(paciente domain.Paciente) error
	// Read devuelve un paciente por su id
	ReadPaciente(id int) (domain.Paciente, error)
	// Update actualiza un paciente
	UpdatePaciente(paciente domain.Paciente) error
	// Delete elimina un paciente
	DeletePaciente(id int) error
	// Exists verifica si un DNI de paciente ya existe
	ExistsDni(dni string) bool
}
