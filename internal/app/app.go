package app

import (
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/controller/http/routes"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
	"github.com/his-vita/patients-service/internal/infrastructure/database"
	"github.com/his-vita/patients-service/internal/infrastructure/httpserver"
	"github.com/his-vita/patients-service/internal/infrastructure/logger"
	"github.com/his-vita/patients-service/internal/infrastructure/sqlstore"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/internal/service"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Env)

	pgContext, err := database.NewPostgresConnect(&cfg.Db)
	if err != nil {
		panic(err)
	}

	sqlStore, err := sqlstore.New(cfg.Sql.Path)
	if err != nil {
		panic(err)
	}

	patientRepository := repository.NewPatientRepository(log, pgContext, sqlStore)
	patientService := service.NewPatientService(patientRepository)
	patientController := v1.NewPatientController(log, patientService)

	httpServer := httpserver.New(cfg.Env, &cfg.Server)

	rg := httpServer.App.Group("/api/v1")

	routes.PatientRoutes(rg, patientController)

	if err := httpServer.Run(&cfg.Server); err != nil {
		panic(err)
	}
}
