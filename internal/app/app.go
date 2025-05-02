package app

import (
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/controller/http/routes"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/internal/service"
	"github.com/his-vita/patients-service/internal/transaction"
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

	patientRepository := repository.NewPatientRepository(pgContext, sqlStore)
	contactRepository := repository.NewContactRepository(pgContext, sqlStore)
	snilsRepository := repository.NewSnilsRepository(pgContext, sqlStore)
	innRepository := repository.NewInnRepository(pgContext, sqlStore)

	patientService := service.NewPatientService(log, patientRepository)
	contactService := service.NewContactService(log, contactRepository)
	snilsService := service.NewSnilsService(log, snilsRepository)
	innService := service.NewInnService(log, innRepository)

	transaction := transaction.NewTransaction(patientService,
		contactService, snilsService, innService, txManager)

	patientController := v1.NewPatientController(patientService, transaction)
	contactController := v1.NewContactController(contactService)
	snilsController := v1.NewSnilsController(snilsService)

	httpServer := httpserver.New(cfg.Env, &cfg.Server)

	rg := httpServer.App.Group("/api/v1")

	routes.PatientRoutes(rg, patientController)
	routes.ContactRoutes(rg, contactController)
	routes.SnilsRoutes(rg, snilsController)

	if err := httpServer.Run(&cfg.Server); err != nil {
		panic(err)
	}
}
