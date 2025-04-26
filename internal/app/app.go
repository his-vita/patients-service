package app

import (
	"github.com/his-vita/patients-service/internal/config"
	contactRepo "github.com/his-vita/patients-service/internal/contact/repository"
	contactServ "github.com/his-vita/patients-service/internal/contact/service"
	"github.com/his-vita/patients-service/internal/controller/http/routes"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
	patientRepo "github.com/his-vita/patients-service/internal/patient/repository"
	patientServ "github.com/his-vita/patients-service/internal/patient/service"
	patientTrans "github.com/his-vita/patients-service/internal/patient/transaction"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/httpserver"
	"github.com/his-vita/patients-service/pkg/logger"
	"github.com/his-vita/patients-service/pkg/sqlstore"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Env)

	pgContext, err := postgres.NewPostgresConnect(&cfg.Db)
	if err != nil {
		panic(err)
	}

	sqlStore, err := sqlstore.New(cfg.Sql.Path)
	if err != nil {
		panic(err)
	}

	txManager := postgres.NewTransactionManager(pgContext)

	patientRepository := patientRepo.New(pgContext, sqlStore)
	contactRepository := contactRepo.New(pgContext, sqlStore)

	patientService := patientServ.New(log, patientRepository)
	contactService := contactServ.New(log, contactRepository)

	patientTransaction := patientTrans.New(patientService, contactService, txManager)

	patientController := v1.NewPatientController(patientService, patientTransaction)
	contactController := v1.NewContactController(contactService)

	httpServer := httpserver.New(cfg.Env, &cfg.Server)

	rg := httpServer.App.Group("/api/v1")

	routes.PatientRoutes(rg, patientController)
	routes.ContactRoutes(rg, contactController)

	if err := httpServer.Run(&cfg.Server); err != nil {
		panic(err)
	}
}
