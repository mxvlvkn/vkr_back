package authhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"wms/internal/modules/auth/dto"
	"wms/internal/modules/auth/service"
	"wms/pkg/utils"
)

type AuthHandler struct {
	authService *authservice.AuthService
}

func NewAuthHandler(authService *authservice.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req authdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.ValidationErrors(
			err,
			map[string]string {
				"Login": "Логин",
				"Password": "Пароль",
			},
		)

		c.Error(fmt.Errorf("Login: Неккоректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		c.Error(fmt.Errorf("Login: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	h.authService.PushToken(c, token)

	response := authdto.LoginResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		c.Error(fmt.Errorf("CheckAuth: Ошибка получения токена: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения токена"})
		return
	}

	_, err = h.authService.ValidateToken(token)
	if err != nil {
		c.Error(fmt.Errorf("CheckAuth: Ошибка валидации токена: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка авторизации"})
		return
	}

	response := authdto.CheckAuthResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.RouterGroup, authService *authservice.AuthService) {
	handler := NewAuthHandler(authService)

	auth := r.Group("/auth")
	auth.POST("/login", handler.Login)
	auth.POST("/check", handler.CheckAuth)
}