package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func InsuranceRoutes(rg *gin.RouterGroup, insuranceCtrl *v1.InsuranceController) {
	g := rg.Group("/insurance")
	{
		g.POST("/", insuranceCtrl.CreateInsurance)
		g.PUT("/:id", insuranceCtrl.UpdateInsurance)
		g.DELETE("/:id", insuranceCtrl.DeleteInsurance)
	}
}
