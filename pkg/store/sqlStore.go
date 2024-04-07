package store

import (
	"database/sql"
	"log"
	"muller-odontologia/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

//--------CRUD ODONTOLOGOS--------

// Create agrega un nuevo odontologo
func (s *sqlStore) CreateOdontologo(odontologo domain.Odontologo) error {
	query := "INSERT INTO odontologos (apellido, nombre, matricula) VALUES (?, ?, ?);"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(odontologo.Apellido, odontologo.Nombre, odontologo.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Read devuelve un odontologo por su id
func (s *sqlStore) ReadOdontologo(id int) (domain.Odontologo, error) {
	var odontologo domain.Odontologo
	query := "SELECT * FROM odontologos WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

// Update actualiza un odontologo
func (s *sqlStore) UpdateOdontologo(odontologo domain.Odontologo) error {
	query := "UPDATE odontologos SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(odontologo.Apellido, odontologo.Nombre, odontologo.Matricula, odontologo.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un odontologo
func (s *sqlStore) DeleteOdontologo(id int) error {
	query := "DELETE FROM odontologos WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// ExistsMatricula verifica si una matricula de odontologo ya existe
func (s *sqlStore) ExistsMatricula(matricula string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM odontologos WHERE matricula = ?;"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

//--------CRUD PACIENTES--------

// Create agrega un nuevo paciente
func (s *sqlStore) CreatePaciente(paciente domain.Paciente) error {
	query := "INSERT INTO pacientes (apellido, nombre, domicilio, dni, alta) VALUES (?, ?, ?, ?, ?);"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(paciente.Apellido, paciente.Nombre, paciente.Domicilio, paciente.Dni, paciente.Alta)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Read devuelve un paciente por su id
func (s *sqlStore) ReadPaciente(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&paciente.Id, &paciente.Apellido, &paciente.Nombre, &paciente.Domicilio, &paciente.Dni, &paciente.Alta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

// Update actualiza un paciente
func (s *sqlStore) UpdatePaciente(paciente domain.Paciente) error {
	query := "UPDATE pacientes SET apellido = ?, nombre = ?, domicilio = ?, dni = ?, alta = ? WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(paciente.Apellido, paciente.Nombre, paciente.Domicilio, paciente.Dni, paciente.Alta, paciente.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un paciente
func (s *sqlStore) DeletePaciente(id int) error {
	query := "DELETE FROM pacientes WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// ExistsDni verifica si un dni de paciente ya existe
func (s *sqlStore) ExistsDni(dni string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM pacientes WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
