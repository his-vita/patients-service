package app

import (
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/controller/http/routes"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
	"github.com/his-vita/patients-service/internal/infrastructure/database"
	"github.com/his-vita/patients-service/internal/infrastructure/httpserver"
	"github.com/his-vita/patients-service/internal/infrastructure/logger"
	"github.com/his-vita/patients-service/internal/infrastructure/sqlstore"
	"github.com/his-vita/patients-service/internal/mapper"
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

	patientMapper := mapper.NewPatientMapper()
	patientRepository := repository.NewPatientRepository(pgContext, sqlStore)
	patientService := service.NewPatientService(log, patientRepository, patientMapper)
	patientController := v1.NewPatientController(patientService)

	contactRepository := repository.NewContactRepository(pgContext, sqlStore)
	contactService := service.NewContactService(log, contactRepository)
	contactController := v1.NewContactController(contactService)

	httpServer := httpserver.New(cfg.Env, &cfg.Server)

	rg := httpServer.App.Group("/api/v1")

	routes.PatientRoutes(rg, patientController)
	routes.ContactRoutes(rg, contactController)

	if err := httpServer.Run(&cfg.Server); err != nil {
		panic(err)
	}
}
