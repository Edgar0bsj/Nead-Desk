package middleware

import (
	"nead-desk/src/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthDefault() gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole, exists := c.Get("user_role")

		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error interno",
			})
			return
		}

		// Só precisa esta autenticado
		if userRole == string(domain.RoleUser) || userRole == string(domain.RoleAdmin) {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Acesso negado!",
		})
		c.Abort()
	}
}
