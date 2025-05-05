package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateUUIDParam(paramName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param(paramName)

		parsedUUID, err := uuid.Parse(idStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid UUID format",
			})
			return
		}

		c.Set(paramName, parsedUUID)
		c.Next()
	}
}
