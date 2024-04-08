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
func validarVaciosOdontologo(odontologo *domain.Odontologo) (bool, error) {
	if odontologo.Apellido == "" {
		return false, errors.New("Debe ingresar un Apellido")
	}
	if odontologo.Nombre == "" {
		return false, errors.New("Debe ingresar un Nombre")
	}
	if odontologo.Matricula == "" {
		return false, errors.New("Debe ingresar una Matricula")
	}
	return true, nil
}

// Post godoc
// @Summary      Creates a new dentist
// @Description  Creates a new dentist in repository
// @Tags         odontologos
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Router       /odontologos [post]
func (h *odontologoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontologo
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
		valid, err := validarVaciosOdontologo(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		o, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, o, "Odontologo creado con exito")
	}
}

// GetByID godoc
// @Summary      Gets a dentist by id
// @Description  Gets a dentist by id from the repository
// @Tags         odontologos
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /odontologos/{id} [get]
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
		web.Success(c, 200, odontologo, "Odontologo encontrado con exito")
	}
}

// Put godoc
// @Summary      Updates a dentist
// @Description  Updates a dentist from the repository
// @Tags         odontologos
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /odontologos/{id} [put]
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
		var odontologo domain.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVaciosOdontologo(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		o, err := h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, o, "Odontologo actualizado con exito")
	}
}

// Delete godoc
// @Summary      Deletes a dentist
// @Description  Deletes a dentist from the repository
// @Tags         odontologos
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /odontologos/{id} [delete]
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
		web.Success(c, 200, nil, "Odontologo eliminado con exito. Se eliminaron tambien sus turnos asociados.")
	}
}

// Patch godoc
// @Summary      Updates selected fields
// @Description  Updates selected fields of a dentist from the repository
// @Tags         odontologos
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Odontologo true "Odontologo"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /odontologos/{id} [patch]
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
		update := domain.Odontologo{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}
		o, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, o, "Odontologo actualizado con exito")
	}
}
