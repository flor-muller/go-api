package store

import (
	"muller-odontologia/internal/domain"
)

type StoreInterface interface {

	//--------ODONTOLOGOS--------

	// Create agrega un nuevo odontologo
	CreateOdontologo(odontologo domain.Odontolgo) error
	// Read devuelve un odontologo por su id
	ReadOdontologo(id int) (domain.Odontolgo, error)
	// Update actualiza un odontologo
	UpdateOdontologo(odontologo domain.Odontolgo) error
	// Delete elimina un odontologo
	DeleteOdontologo(id int) error
	// Exists verifica si una matricula de odontologo ya existe
	Exists(matricula string) bool

	// Update actualiza un odontologo por alguno de sus campos
	//UpdateCampoOdontolgo(odontologo domain.Odontolgo) error
	//GetAll() ([]domain.Odontologo, error)

}
