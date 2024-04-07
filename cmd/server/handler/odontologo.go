package handler

import (
	"errors"
	"muller-odontologia/internal/domain"
	"muller-odontologia/internal/odontologo"
	"muller-odontologia/pkg/web"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// Funcion para validar que los campos no esten vacios
func validarVacios(odontologo *domain.Odontolgo) (bool, error) {
	if odontologo.Apellido == "" || odontologo.Nombre == "" || odontologo.Matricula == "" {
		return false, errors.New("Los campos no pueden estar vacios")
	}
	return true, nil
}

// Post crea un nuevo odontologo
func (h *odontologoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontolgo
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVacios(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		o, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, o)
	}
}

// Get devuelve un odontologo por id
func (h *odontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El id ingresado es invalido"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, odontologo)
	}
}

// Put actualiza un odontologo
func (h *odontologoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var odontologo domain.Odontolgo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVacios(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		o, err := h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, o)
	}
}

// Delete elimina un odontologo
func (h *odontologoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Patch actualiza un odontologo o alguno de sus campos
func (h *odontologoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Apellido  string `json:"apellido,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
		Matricula string `json:"matricula,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		update := domain.Odontolgo{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}
		o, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, o)
	}
}
