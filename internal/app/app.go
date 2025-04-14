package app

import (
	"log/slog"

	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/controller"
	"github.com/his-vita/patients-service/internal/database"
	httpserver "github.com/his-vita/patients-service/internal/http-server"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/internal/routes"
	"github.com/his-vita/patients-service/internal/service"
)

type App struct {
	httpServer *httpserver.HttpServer
	pgContext  *database.PgContext
	log        *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *App {
	pgContext := database.NewPostgresConnect(&cfg.Db)

	patientRepository := repository.NewPatientRepository(log, pgContext, cfg.Db.SqlPath)
	patientService := service.NewPatientService(patientRepository)
	patientController := controller.NewPatientController(log, patientService)

	httpServer := httpserver.New(cfg.Env, &cfg.Server)

	rg := httpServer.RouterGroup().Group("/api/v1")

	routes.PatientRoutes(rg, patientController)

	return &App{
		httpServer: httpServer,
		pgContext:  pgContext,
		log:        log,
	}
}

func (a *App) MustRun() {
	a.httpServer.MustRun()
}

func (a *App) Close() {
	a.pgContext.Pool.Close()
}
