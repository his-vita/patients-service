package main

import (
	"log/slog"

	"github.com/his-vita/patients-service/internal/app"
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)
	log.Info("config and logger loaded", slog.Any("config", cfg))

	application := app.New(cfg, log)
	defer application.Close()

	application.MustRun()
}
