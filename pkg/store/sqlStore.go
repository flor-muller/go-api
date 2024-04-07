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

//--------ODONTOLOGOS--------

// Create agrega un nuevo odontologo
func (s *sqlStore) CreateOdontologo(odontologo domain.Odontolgo) error {
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
func (s *sqlStore) ReadOdontologo(id int) (domain.Odontolgo, error) {
	var odontologo domain.Odontolgo
	query := "SELECT * FROM odontologos WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&odontologo.Id, &odontologo.Apellido, &odontologo.Nombre, &odontologo.Matricula)
	if err != nil {
		return domain.Odontolgo{}, err
	}
	return odontologo, nil
}

// Update actualiza un odontologo
func (s *sqlStore) UpdateOdontologo(odontologo domain.Odontolgo) error {
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

// Exists verifica si una matricula de odontologo ya existe
func (s *sqlStore) Exists(matricula string) bool {
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
