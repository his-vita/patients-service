package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller/http/middleware"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func SnilsRoutes(rg *gin.RouterGroup, snilsCtrl *v1.SnilsController) {
	cr := rg.Group("/snils")
	{
		cr.PUT("/:id", middleware.ValidateUUIDParam("id"), snilsCtrl.UpdateSnils)
	}
}
