package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller/http/middleware"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func InsuranceRoutes(rg *gin.RouterGroup, insuranceCtrl *v1.InsuranceController) {
	g := rg.Group("/insurance")
	{
		g.POST("/", insuranceCtrl.CreateInsurance)
		g.PUT("/:id", middleware.ValidateUUIDParam("id"), insuranceCtrl.UpdateInsurance)
		g.DELETE("/:id", middleware.ValidateUUIDParam("id"), insuranceCtrl.DeleteInsurance)
	}
}
