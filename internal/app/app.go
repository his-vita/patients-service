package app

import (
	httpserver "github.com/his-vita/patients-service/internal/app/http-server"
	"github.com/his-vita/patients-service/internal/controller"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/internal/routes"
	"github.com/his-vita/patients-service/internal/service"
)

type App struct {
	HttpServer *httpserver.HttpServer
}

func New() *App {
	patientRepository := repository.NewPatientRepository()
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(patientService)

	httpServer := httpserver.New()

	rg := httpServer.RouterGroup().Group("/api/v1")

	routes.PatientRoutes(rg, patientController)

	return &App{
		HttpServer: httpServer,
	}
}
