package app

import (
	httpserver "github.com/his-vita/patients-service/internal/app/http-server"
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/controller"
	"github.com/his-vita/patients-service/internal/database"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/internal/routes"
	"github.com/his-vita/patients-service/internal/service"
)

type App struct {
	HttpServer *httpserver.HttpServer
}

func New(dbCfg *config.Db) *App {
	pg := database.NewPostgresConnect(dbCfg)
	defer pg.CloseConnect()

	patientRepository := repository.NewPatientRepository(pg.PgCon)
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(patientService)

	httpServer := httpserver.New()

	rg := httpServer.RouterGroup().Group("/api/v1")

	routes.PatientRoutes(rg, patientController)

	return &App{
		HttpServer: httpServer,
	}
}
