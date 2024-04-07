package odontologo

import (
	"muller-odontologia/internal/domain"
)

type Service interface {
	// Create agrega un nuevo odontologo
	Create(odontologo domain.Odontologo) (domain.Odontologo, error)
	// GetByID devuelve un odontologo por id
	GetByID(id int) (domain.Odontologo, error)
	// Update actualiza un odontologo
	Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error)
	// Delete elimina un odontologo
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(odontologo domain.Odontologo) (domain.Odontologo, error) {
	odontologo, err := s.r.Create(odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *service) Update(id int, u domain.Odontologo) (domain.Odontologo, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	if u.Apellido != "" {
		o.Apellido = u.Apellido
	}
	if u.Nombre != "" {
		o.Nombre = u.Nombre
	}
	if u.Matricula != "" {
		o.Matricula = u.Matricula
	}

	o, err = s.r.Update(id, o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
