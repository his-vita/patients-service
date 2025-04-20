package main

import (
	"github.com/his-vita/patients-service/internal/app"
	"github.com/his-vita/patients-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
