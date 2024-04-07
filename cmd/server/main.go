package main

import (
	"database/sql"
	"log"
	"muller-odontologia/cmd/server/handler"
	"muller-odontologia/internal/odontologo"
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
	repo := odontologo.NewRepository(storage)
	service := odontologo.NewService(repo)
	odontologoHandler := handler.NewOdontologoHandler(service)

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

	r.Run(":8080")
}
