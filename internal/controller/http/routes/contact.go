package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/his-vita/patients-service/internal/controller/http/v1"
)

func ContactRoutes(rg *gin.RouterGroup, contactCtrl *v1.ContactController) {
	cr := rg.Group("/contacts")
	{
		cr.PUT("/", contactCtrl.UpdateContact)
	}
}
