package middleware

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
    "wms/internal/modules/auth/service"
)

func AuthMiddleware(authService *authservice.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil {
			c.Error(fmt.Errorf("AuthMiddleware: Ошибка получения токена."))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения токена"})
			return
		}

		claims, err := authService.ValidateToken(token)
		if err != nil {
			c.Error(fmt.Errorf("AuthMiddleware: Ошибка валидации токена: %w", err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка авторизации"})
			return
		}

		c.Set("claims", claims) //!c.Get("claims")
		c.Next()
	}
}