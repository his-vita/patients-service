package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func SnilsRoutes(rg *gin.RouterGroup, snilsCtrl *v1.SnilsController) {
	cr := rg.Group("/snils")
	{
		cr.PUT("/", snilsCtrl.UpdateSnils)
	}
}
