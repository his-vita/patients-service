package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller/http/middleware"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func DocumentRoutes(rg *gin.RouterGroup, documentCtrl *v1.DocumentController) {
	g := rg.Group("/documents")
	{
		g.POST("/", documentCtrl.CreateDocument)
		g.PUT("/:id", middleware.ValidateUUIDParam("id"), documentCtrl.UpdateDocument)
		g.DELETE("/:id", middleware.ValidateUUIDParam("id"), documentCtrl.DeleteDocument)
	}
}
