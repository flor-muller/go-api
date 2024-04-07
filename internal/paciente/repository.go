package paciente

import (
	"errors"
	"muller-odontologia/internal/domain"
	"muller-odontologia/pkg/store"
)

type Repository interface {
	// Create agrega un nuevo paciente
	Create(paciente domain.Paciente) (domain.Paciente, error)
	// GetByID devuelve un paciente por id
	GetByID(id int) (domain.Paciente, error)
	// Update actualiza un paciente
	Update(id int, paciente domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(paciente domain.Paciente) (domain.Paciente, error) {
	if r.storage.ExistsDni(paciente.Dni) {
		return domain.Paciente{}, errors.New("El DNI ingresado ya existe.")
	}
	err := r.storage.CreatePaciente(paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Error al crear paciente.")
	}
	return paciente, nil
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.ReadPaciente(id)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente no encontrado.")
	}
	return paciente, nil

}

func (r *repository) Update(id int, paciente domain.Paciente) (domain.Paciente, error) {
	err := r.storage.UpdatePaciente(paciente)
	if err != nil {
		return domain.Paciente{}, errors.New("Error al actualizar paciente.")
	}
	return paciente, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	err = r.storage.DeletePaciente(id)
	if err != nil {
		return err
	}
	return nil
}
