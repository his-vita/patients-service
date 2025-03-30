package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/his-vita/patients-service/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("starting application")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	if env == envLocal {
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	} else if env == envProd {
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
