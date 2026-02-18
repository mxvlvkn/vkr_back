package authhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/auth/dto"
	thisservice "wms/internal/modules/auth/service"
	"wms/pkg/utils"
)

type Handler struct {
	service *thisservice.Service
}

func NewHandler(service *thisservice.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req thisdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.ValidationErrors(err)

		c.Error(fmt.Errorf("Login: Неккоректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	loginData, err := h.service.Login(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		c.Error(fmt.Errorf("Login: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	h.service.PushToken(c, loginData["token"])

	response := thisdto.LoginResponse {
		Status: true,
		Login: loginData["login"],
		Role: loginData["role"],
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		c.Error(fmt.Errorf("CheckAuth: Ошибка получения токена: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения токена"})
		return
	}

	_, err = h.service.ValidateToken(token)
	if err != nil {
		c.Error(fmt.Errorf("CheckAuth: Ошибка валидации токена: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка авторизации"})
		return
	}

	response := thisdto.CheckAuthResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.RouterGroup, service *thisservice.Service) {
	handler := NewHandler(service)

	api := r.Group("/auth")
	api.POST("/login", handler.Login)
	api.GET("/check", handler.CheckAuth)
}