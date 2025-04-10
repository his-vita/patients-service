package main

import (
	"fmt"
	"log/slog"

	"github.com/his-vita/patients-service/internal/app"
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)
	log.Info("config and logger loaded", slog.Any("config", cfg))

	application := app.New(&cfg.Db, log)
	defer application.PgContext.Close()

	application.HttpServer.MustRun(fmt.Sprintf(":%v", cfg.Server.Port))
}
