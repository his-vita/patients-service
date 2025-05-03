package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller/http/middleware"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func InnRoutes(rg *gin.RouterGroup, innCtrl *v1.InnController) {
	cr := rg.Group("/inn")
	{
		cr.PUT("/:id", middleware.ValidateUUIDParam("id"), innCtrl.UpdateInn)
	}
}
