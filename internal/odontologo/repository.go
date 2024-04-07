package odontologo

import (
	"errors"
	"muller-odontologia/internal/domain"
	"muller-odontologia/pkg/store"
)

type Repository interface {
	// Create agrega un nuevo odontologo
	Create(odontologo domain.Odontolgo) (domain.Odontolgo, error)
	// GetByID devuelve un odontologo por id
	GetByID(id int) (domain.Odontolgo, error)
	// Update actualiza un odontologo
	Update(id int, odontologo domain.Odontolgo) (domain.Odontolgo, error)
	// Delete elimina un odontologo
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(odontologo domain.Odontolgo) (domain.Odontolgo, error) {
	if r.storage.Exists(odontologo.Matricula) {
		return domain.Odontolgo{}, errors.New("La matricula ingresada ya existe.")
	}
	err := r.storage.CreateOdontologo(odontologo)
	if err != nil {
		return domain.Odontolgo{}, errors.New("Error al crear odontologo.")
	}
	return odontologo, nil
}

func (r *repository) GetByID(id int) (domain.Odontolgo, error) {
	odontologo, err := r.storage.ReadOdontologo(id)
	if err != nil {
		return domain.Odontolgo{}, errors.New("Odontologo no encontrado.")
	}
	return odontologo, nil

}

func (r *repository) Update(id int, odontologo domain.Odontolgo) (domain.Odontolgo, error) {
	err := r.storage.UpdateOdontologo(odontologo)
	if err != nil {
		return domain.Odontolgo{}, errors.New("Error al actualizar odontologo.")
	}
	return odontologo, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeleteOdontologo(id)
	if err != nil {
		return err
	}
	return nil
}
