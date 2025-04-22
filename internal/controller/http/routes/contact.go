package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller/http/middleware"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func ContactRoutes(rg *gin.RouterGroup, contactCtrl *v1.ContactController) {
	cr := rg.Group("/contacts")
	{
		cr.GET("/:id", middleware.ValidateUUIDParam("id"), contactCtrl.GetContactsByPatientId)
		cr.POST("/", contactCtrl.CreateContact)
		cr.PUT("/", contactCtrl.UpdateContact)
		cr.DELETE("/:id", middleware.ValidateUUIDParam("id"), contactCtrl.DeleteContact)
	}
}
