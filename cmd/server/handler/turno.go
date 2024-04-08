package handler

import (
	"errors"
	"muller-odontologia/internal/domain"
	"muller-odontologia/internal/turno"
	"muller-odontologia/pkg/web"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.Service
}

func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// Funcion para validar que los campos no esten vacios
func validarVaciosTurnos(turno *domain.Turno) (bool, error) {
	if turno.IdPaciente <= 0 {
		return false, errors.New("Debe ingresar un IdPaciente valido")
	}
	if turno.IdOdontologo <= 0 {
		return false, errors.New("Debe ingresar un IdOdontologo valido")
	}
	if turno.Fecha == "" {
		return false, errors.New("Debe ingresar una Fecha")
	}
	if turno.Hora == "" {
		return false, errors.New("Debe ingresar una Hora")
	}
	if turno.Descripcion == "" {
		return false, errors.New("Debe ingresar una Descripcion")
	}
	return true, nil
}

// Funcion para validar que los campos no esten vacios
func validarVaciosTurnosDM(turnoDM *domain.TurnoDM) (bool, error) {
	if turnoDM.Dni == "" {
		return false, errors.New("Debe ingresar un DNI")
	}
	if turnoDM.Matricula == "" {
		return false, errors.New("Debe ingresar una Matricula")
	}
	if turnoDM.Fecha == "" {
		return false, errors.New("Debe ingresar una Fecha")
	}
	if turnoDM.Hora == "" {
		return false, errors.New("Debe ingresar una Hora")
	}
	if turnoDM.Descripcion == "" {
		return false, errors.New("Debe ingresar una Descripcion")
	}
	return true, nil
}

// validarFecha verifica que la fecha sea valida
func validarFecha(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("Fecha de turno invalida. Debe ingresar en formato: dd/mm/aaaa")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("Fecha de turno invalida. Se deben ingresar numeros")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("Fecha de turno invalida. Revise los valores asignados a dia, mes y/o año (dd/mm/aaa)")
	}
	return true, nil
}

// validarHora verifica que la hora sea valida
func validarHora(exp string) (bool, error) {
	hour := strings.Split(exp, ":")
	list := []int{}
	if len(hour) != 2 {
		return false, errors.New("Hora de turno invalida. Debe ingresar en formato: hh:mm")
	}
	for value := range hour {
		number, err := strconv.Atoi(hour[value])
		if err != nil {
			return false, errors.New("Hora de turno invalida. Se deben ingresar numeros")
		}
		list = append(list, number)
	}
	condition := (list[0] < 7 || list[0] > 19) || (list[1] < 0 || list[1] > 59)
	if condition {
		return false, errors.New("Hora de turno invalida. Revise los valores, la hora debe estar entre las 07:00 y las 20:00")
	}
	return true, nil
}

// Post godoc
// @Summary      Creates a appointment
// @Description  Creates a appointment in repository
// @Tags         turnos
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Turno true "Turno"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Router       /turnos [post]
func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVaciosTurnos(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarFecha(turno.Fecha)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarHora(turno.Hora)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		t, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, t, "Turno creado con exito")
	}
}

// GetByID godoc
// @Summary      Gets an appointment by id
// @Description  Gets an appointment by id from the repository
// @Tags         turnos
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /turnos/{id} [get]
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("El id ingresado es invalido"))
			return
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, turno, "Turno encontrado con exito")
	}
}

// Put godoc
// @Summary      Updates a appointment
// @Description  Updates a appointment from the repository
// @Tags         turnos
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Turno true "Turno"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /turnos/{id} [put]
func (h *turnoHandler) Put() gin.HandlerFunc {
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
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVaciosTurnos(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarFecha(turno.Fecha)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarHora(turno.Hora)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		t, err := h.s.Update(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t, "Turno actualizado con exito")
	}
}

// Delete godoc
// @Summary      Deletes an appoointmet
// @Description  Deletes an appoointmet from the repository
// @Tags         turnos
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /turnos/{id} [delete]
func (h *turnoHandler) Delete() gin.HandlerFunc {
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
		web.Success(c, 200, nil, "Turno eliminado con exito.")
	}
}

// Patch godoc
// @Summary      Updates selected fields
// @Description  Updates selected fields of an appointment from the repository
// @Tags         turnos
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Turno true "Turno"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Failure      409 {object}  web.errorResponse
// @Router       /turnos/{id} [patch]
func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		IdPaciente   int    `json:"id_paciente,omitempty"`
		IdOdontologo int    `json:"id_odontologo,omitempty"`
		Fecha        string `json:"fecha,omitempty"`
		Hora         string `json:"hora,omitempty"`
		Descripcion  string `json:"descripcion,omitempty"`
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
		update := domain.Turno{
			IdPaciente:   r.IdPaciente,
			IdOdontologo: r.IdOdontologo,
			Fecha:        r.Fecha,
			Hora:         r.Hora,
			Descripcion:  r.Descripcion,
		}
		if update.Fecha != "" {
			valid, err := validarFecha(update.Fecha)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		if update.Hora != "" {
			valid, err := validarHora(update.Hora)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		t, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t, "Turno actualizado con exito")
	}
}

// PostByDM godoc
// @Summary      Creates a appointment by patient's dni and dentist's license
// @Description  Creates a appointment in repository
// @Tags         turnos
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.TurnoDM true "TurnoDM"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      401 {object}  web.errorResponse
// @Router       /turnos/dm [post]
func (h *turnoHandler) PostByDM() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turnoDM domain.TurnoDM
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token no encontrado"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Token invalido"))
			return
		}
		err := c.ShouldBindJSON(&turnoDM)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON invalido"))
			return
		}
		valid, err := validarVaciosTurnosDM(&turnoDM)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarFecha(turnoDM.Fecha)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validarHora(turnoDM.Hora)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		t, err := h.s.CreateTurnoDniMatricula(turnoDM)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, t, "Turno creado con exito para el DNI de paciente y matricula de odontologo informados.")
	}
}

// GetByDni godoc
// @Summary      Gets an appointment by patient's dni
// @Description  Gets an appointment by patient's dni from the repository
// @Tags         turnos
// @Produce      json
// @Param        dni query string true "DNI"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /turnos/dni [get]
func (h *turnoHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Query("dni")
		if dni == "" {
			web.Failure(c, 400, errors.New("Debe ingresar un DNI"))
			return
		}
		listaTurnos, err := h.s.GetByDni(dni)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		if len(listaTurnos) <= 0 {
			web.Failure(c, 400, errors.New("No existen turnos para el DNI ingresado"))
			return
		}
		web.Success(c, 200, listaTurnos, "Turno/s encontrado/s con exito")
	}
}
