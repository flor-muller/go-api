package main

import (
	"database/sql"
	"log"
	"muller-odontologia/cmd/server/handler"
	"muller-odontologia/internal/odontologo"
	"muller-odontologia/internal/paciente"
	"muller-odontologia/pkg/middleware"
	"muller-odontologia/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos_odontologia")
	if err != nil {
		log.Fatal(err)
	}

	storage := store.NewSqlStore(bd)

	repoOdontologo := odontologo.NewRepository(storage)
	serviceOdontologo := odontologo.NewService(repoOdontologo)
	odontologoHandler := handler.NewOdontologoHandler(serviceOdontologo)

	repoPaciente := paciente.NewRepository(storage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	r := gin.Default()

	//--------ODONTOLOGOS--------
	odontologos := r.Group("/odontologos")
	{
		odontologos.POST("", middleware.Authentication(), odontologoHandler.Post())
		odontologos.GET(":id", odontologoHandler.GetByID())
		odontologos.PUT(":id", middleware.Authentication(), odontologoHandler.Put())
		odontologos.DELETE(":id", middleware.Authentication(), odontologoHandler.Delete())
		odontologos.PATCH(":id", middleware.Authentication(), odontologoHandler.Patch())
	}

	//--------PACIENTES--------
	pacientes := r.Group("/pacientes")
	{
		pacientes.POST("", middleware.Authentication(), pacienteHandler.Post())
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.PUT(":id", middleware.Authentication(), pacienteHandler.Put())
		pacientes.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())
		pacientes.PATCH(":id", middleware.Authentication(), pacienteHandler.Patch())
	}

	r.Run(":8080")
}
