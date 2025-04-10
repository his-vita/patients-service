package app

import (
	"log/slog"

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
	PgContext  *database.PgContext
	log        *slog.Logger
}

func New(dbCfg *config.Db, log *slog.Logger) *App {
	pgContext := database.NewPostgresConnect(dbCfg)

	patientRepository := repository.NewPatientRepository(pgContext, dbCfg.SqlPath)
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(patientService)

	httpServer := httpserver.New()

	rg := httpServer.RouterGroup().Group("/api/v1")

	routes.PatientRoutes(rg, patientController)

	return &App{
		HttpServer: httpServer,
		PgContext:  pgContext,
		log:        log,
	}
}
