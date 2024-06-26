package handler

import (
	"errors"
	"muller-odontologia/internal/domain"
	"muller-odontologia/internal/paciente"
	"muller-odontologia/pkg/web"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// Funcion para validar que los campos no esten vacios
func validarVaciosPaciente(paciente *domain.Paciente) (bool, error) {
	if paciente.Apellido == "" {
		return false, errors.New("Debe ingresar un Apellido")
	}
	if paciente.Nombre == "" {
		return false, errors.New("Debe ingresar un Nombre")
	}
	if paciente.Domicilio == "" {
		return false, errors.New("Debe ingresar un Domicilio")
	}
	if paciente.Dni == "" {
		return false, errors.New("Debe ingresar un Dni")
	}
	if paciente.Alta == "" {
		return false, errors.New("Debe ingresar una Fecha de Alta")
	}
	return true, nil
}

// validarAlta verifica que la fecha de alta sea valida
func validarAlta(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("Fecha de alta invalida. Debe ingresar en formato: dd/mm/aaaa")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("Fecha de alta invalida. Se deben ingresar numeros")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("Fecha de alta invalida. Revise los valores asignados a dia, mes y/o año (dd/mm/aaa)")
	}
	return true, nil
}

// Post godoc
// @Summary      Creates a new patient
// @Description  Creates a new patient in repository
// @Tags         pacientes
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Paciente true "Paciente"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /pacientes [post]
func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVaciosPaciente(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarAlta(paciente.Alta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p, "Paciente creado con exito")
	}
}

// GetByID godoc
// @Summary      Gets a patient by id
// @Description  Gets a patient by id from the repository
// @Tags         pacientes
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /pacientes/{id} [get]
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El id ingresado es invalido"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, paciente, "Paciente encontrado con exito")
	}
}

// Put godoc
// @Summary      Updates a patient
// @Description  Updates a patient from the repository
// @Tags         pacientes
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Paciente true "Paciente"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /pacientes/{id} [put]
func (h *pacienteHandler) Put() gin.HandlerFunc {
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
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVaciosPaciente(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarAlta(paciente.Alta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "Paciente actualizado con exito")
	}
}

// Delete godoc
// @Summary      Deletes a patient
// @Description  Deletes a patient from the repository
// @Tags         pacientes
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /pacientes/{id} [delete]
func (h *pacienteHandler) Delete() gin.HandlerFunc {
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
		web.Success(c, 200, nil, "Paciente eliminado con exito. Se eliminaron tambien sus turnos asociados.")
	}
}

// Patch godoc
// @Summary      Updates selected fields
// @Description  Updates selected fields of a patient from the repository
// @Tags         pacientes
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Paciente true "Paciente"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /pacientes/{id} [patch]
func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Apellido  string `json:"apellido,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
		Domicilio string `json:"domicilio,omitempty"`
		Dni       string `json:"dni,omitempty"`
		Alta      string `json:"alta,omitempty"`
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
		update := domain.Paciente{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Domicilio: r.Domicilio,
			Dni:       r.Dni,
			Alta:      r.Alta,
		}
		if update.Alta != "" {
			valid, err := validarAlta(update.Alta)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p, "Paciente actualizado con exito")
	}
}
