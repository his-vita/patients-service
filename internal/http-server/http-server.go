package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/config"
)

type HttpServer struct {
	server *gin.Engine
	config *config.Server
}

func New(env string, cfg *config.Server) *HttpServer {
	if env == config.EnvLocal {
		gin.SetMode(gin.DebugMode)
	} else if env == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	return &HttpServer{
		server: r,
		config: cfg,
	}
}

func (h *HttpServer) MustRun() {
	addr := fmt.Sprintf("%s:%v", h.config.Host, h.config.Port)

	if err := h.server.Run(addr); err != nil {
		panic(err)
	}
}

func (h *HttpServer) RouterGroup() *gin.RouterGroup {
	return &h.server.RouterGroup
}
