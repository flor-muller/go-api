package paciente

import (
	"muller-odontologia/internal/domain"
)

type Service interface {
	// Create agrega un nuevo paciente
	Create(paciente domain.Paciente) (domain.Paciente, error)
	// GetByID devuelve un paciente por id
	GetByID(id int) (domain.Paciente, error)
	// Update actualiza un paciente
	Update(id int, paciente domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(paciente domain.Paciente) (domain.Paciente, error) {
	paciente, err := s.r.Create(paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *service) GetByID(id int) (domain.Paciente, error) {
	paciente, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Domicilio != "" {
		p.Domicilio = u.Domicilio
	}
	if u.Dni != "" {
		p.Dni = u.Dni
	}
	if u.Alta != "" {
		p.Alta = u.Alta
	}

	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
