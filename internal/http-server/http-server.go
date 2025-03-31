package httpserver

import "github.com/gin-gonic/gin"

type HttpServer struct {
	server *gin.Engine
}

func New() *HttpServer {
	return &HttpServer{}
}

func (h *HttpServer) MustRun() {
	h.server.Run()
}

func (h *HttpServer) routesSetup() {
	h.server.Group("v1")
}
