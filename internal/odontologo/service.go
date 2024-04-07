package odontologo

import (
	"muller-odontologia/internal/domain"
)

type Service interface {
	// Create agrega un nuevo odontologo
	Create(odontologo domain.Odontolgo) (domain.Odontolgo, error)
	// GetByID devuelve un odontologo por id
	GetByID(id int) (domain.Odontolgo, error)
	// Update actualiza un odontologo
	Update(id int, odontologo domain.Odontolgo) (domain.Odontolgo, error)
	// Delete elimina un odontologo
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(odontologo domain.Odontolgo) (domain.Odontolgo, error) {
	odontologo, err := s.r.Create(odontologo)
	if err != nil {
		return domain.Odontolgo{}, err
	}
	return odontologo, nil
}

func (s *service) GetByID(id int) (domain.Odontolgo, error) {
	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontolgo{}, err
	}
	return odontologo, nil
}

func (s *service) Update(id int, u domain.Odontolgo) (domain.Odontolgo, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontolgo{}, err
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
		return domain.Odontolgo{}, err
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
