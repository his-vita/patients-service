package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/his-vita/patients-service/internal/app"
	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	application := app.New(&cfg.Db)

	go application.HttpServer.MustRun(fmt.Sprintf(":%v", cfg.Server.Port))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))
}
