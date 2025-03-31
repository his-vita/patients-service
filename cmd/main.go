package main

import (
	"fmt"

	"github.com/his-vita/patients-service/internal/config"
	"github.com/his-vita/patients-service/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := logger.New(cfg.Env)
	log.Info("starting application")
}
