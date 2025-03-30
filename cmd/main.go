package main

import (
	"fmt"

	"github.com/his-vita/patients-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
