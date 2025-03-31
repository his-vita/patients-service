package httpserver

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *gin.Engine
}

func New() *HttpServer {
	r := gin.Default()

	return &HttpServer{
		server: r,
	}
}

func (a *HttpServer) MustRun(addr string) {
	if err := a.server.Run(addr); err != nil {
		panic(err)
	}
}

func (h *HttpServer) RouterGroup() *gin.RouterGroup {
	return &h.server.RouterGroup
}
