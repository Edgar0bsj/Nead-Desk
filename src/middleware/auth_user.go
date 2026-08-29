package middleware

import (
	"nead-desk/src/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func auth_user() gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole, exists := c.Get("user_role")

		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error interno",
			})
			return
		}

		// Somente acesso User
		if userRole == string(domain.RoleUser) {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Acesso negado!",
		})
		c.Abort()
	}
}
