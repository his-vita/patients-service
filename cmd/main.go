package main

import (
	"fmt"

	"github.com/his-vita/patients-service/internal/app"
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	application := app.New(&cfg.Db, log)
	defer application.PgContext.Close()

	application.HttpServer.MustRun(fmt.Sprintf(":%v", cfg.Server.Port))
}
