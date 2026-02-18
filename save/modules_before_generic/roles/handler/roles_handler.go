package roleshandler

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	"wms/internal/modules/roles/dto"
	"wms/internal/modules/roles/service"
)

type RolesHandler struct {
	rolesService *rolesservice.RolesService
}

func NewRolesHandler(rolesService *rolesservice.RolesService) *RolesHandler {
	return &RolesHandler{
		rolesService: rolesService,
	}
}

func (h *RolesHandler) GetAll(c *gin.Context) {
	roles, err := h.rolesService.GetAll(c)
	if err != nil {
		c.Error(fmt.Errorf("GetAll: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения должностей"})
		return
	}

	response := rolesdto.GetAllResponse {
		Roles: *roles,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.RouterGroup, rolesService *rolesservice.RolesService) {
	handler := NewRolesHandler(rolesService)

	auth := r.Group("/roles")
	{
		auth.POST("/getall", handler.GetAll)
	}
}