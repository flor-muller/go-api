package turno

import (
	"muller-odontologia/internal/domain"
)

type Service interface {
	// Create agrega un nuevo turno
	Create(turno domain.Turno) (domain.Turno, error)
	// GetByID devuelve un turno por id
	GetByID(id int) (domain.Turno, error)
	// Update actualiza un turno
	Update(id int, turno domain.Turno) (domain.Turno, error)
	// Delete elimina un turno
	Delete(id int) error
	// CreateTurnoDniMatricula agrega un nuevo turno por DNI del paciente y matrícula del odontologo
	CreateTurnoDniMatricula(turnoDM domain.TurnoDM) (domain.TurnoDM, error)
	// ReadTurnoDni devuelve turno por DNI del paciente
	GetByDni(dni string) ([]domain.TurnoDetalle, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(turno domain.Turno) (domain.Turno, error) {
	turno, err := s.r.Create(turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *service) GetByID(id int) (domain.Turno, error) {
	turno, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *service) Update(id int, u domain.Turno) (domain.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if u.IdPaciente > 0 {
		t.IdPaciente = u.IdPaciente
	}
	if u.IdOdontologo > 0 {
		t.IdOdontologo = u.IdOdontologo
	}
	if u.Fecha != "" {
		t.Fecha = u.Fecha
	}
	if u.Hora != "" {
		t.Hora = u.Hora
	}
	if u.Descripcion != "" {
		t.Descripcion = u.Descripcion
	}

	t, err = s.r.Update(id, t)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateTurnoDniMatricula(turnoDM domain.TurnoDM) (domain.TurnoDM, error) {
	turnoDM, err := s.r.CreateTurnoDniMatricula(turnoDM)
	if err != nil {
		return domain.TurnoDM{}, err
	}
	return turnoDM, nil
}

func (s *service) GetByDni(dni string) ([]domain.TurnoDetalle, error) {
	listaTurnos, err := s.r.GetByDni(dni)
	if err != nil {
		return []domain.TurnoDetalle{}, err
	}
	return listaTurnos, nil
}
