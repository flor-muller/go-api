package turno

import (
	"errors"
	"muller-odontologia/internal/domain"
	"muller-odontologia/pkg/store"
)

type Repository interface {
	// Create agrega un nuevo turno
	Create(turno domain.Turno) (domain.Turno, error)
	// GetByID devuelve un turno por id
	GetByID(id int) (domain.Turno, error)
	// Update actualiza un turno
	Update(id int, paciente domain.Turno) (domain.Turno, error)
	// Delete elimina un turno
	Delete(id int) error
	// CreateTurnoDniMatricula agrega un nuevo turno por DNI del paciente y matrícula del odontologo
	CreateTurnoDniMatricula(turnoDM domain.TurnoDM) (domain.TurnoDM, error)
	// ReadTurnoDni devuelve turno por DNI del paciente
	GetByDni(dni string) ([]domain.TurnoDetalle, error)
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) Create(turno domain.Turno) (domain.Turno, error) {
	err := r.storage.CreateTurno(turno)
	if err != nil {
		return domain.Turno{}, errors.New("Error al crear turno.")
	}
	return turno, nil
}

func (r *repository) GetByID(id int) (domain.Turno, error) {
	turno, err := r.storage.ReadTurno(id)
	if err != nil {
		return domain.Turno{}, errors.New("Turno no encontrado.")
	}
	return turno, nil

}

func (r *repository) Update(id int, turno domain.Turno) (domain.Turno, error) {
	err := r.storage.UpdateTurno(turno)
	if err != nil {
		return domain.Turno{}, errors.New("Error al actualizar turno.")
	}
	return turno, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	err = r.storage.DeleteTurno(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateTurnoDniMatricula(turnoDM domain.TurnoDM) (domain.TurnoDM, error) {
	err := r.storage.CreateTurnoDniMatricula(turnoDM)
	if err != nil {
		return domain.TurnoDM{}, errors.New("Error al crear turno.")
	}
	return turnoDM, nil
}

func (r *repository) GetByDni(dni string) ([]domain.TurnoDetalle, error) {
	listaTurnos, err := r.storage.ReadTurnoDni(dni)
	if err != nil {
		return []domain.TurnoDetalle{}, errors.New("No existen turnos para el DNI ingresado.")
	}
	return listaTurnos, nil

}
