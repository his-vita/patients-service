package logger

import (
	"log/slog"
	"os"

	"github.com/his-vita/patients-service/internal/config"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	if env == config.EnvLocal {
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	} else if env == config.EnvProd {
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
