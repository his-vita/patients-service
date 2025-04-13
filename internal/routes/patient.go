package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller"
	"github.com/his-vita/patients-service/internal/middleware"
)

func PatientRoutes(rg *gin.RouterGroup, patientCtrl *controller.PatientController) {
	pg := rg.Group("/patients")
	{
		pg.GET("/:id", middleware.ValidateUUIDParam("id"), patientCtrl.GetPatient)
		pg.POST("/", patientCtrl.CreatePatient)
		pg.PUT("/:id", patientCtrl.UpdatePatient)
		pg.PATCH("/mark_deleted/:id", middleware.ValidateUUIDParam("id"), patientCtrl.MarkPatientAsDeleted)
		pg.PATCH("/unmark_deleted/:id", middleware.ValidateUUIDParam("id"), patientCtrl.UnMarkPatientAsDeleted)
	}
}
