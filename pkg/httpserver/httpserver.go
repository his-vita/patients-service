package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/config"
)

type Server struct {
	App *gin.Engine
}

func New(env string, cfg *config.Server) *Server {
	if env == config.EnvLocal {
		gin.SetMode(gin.DebugMode)
	} else if env == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	return &Server{
		App: r,
	}
}

func (s *Server) Run(cfg *config.Server) error {
	addr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)

	if err := s.App.Run(addr); err != nil {
		return fmt.Errorf("error http server: %s", err)
	}

	return nil
}
