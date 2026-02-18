package unitshandler

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"wms/internal/modules/units/dto"
	"wms/internal/modules/units/service"
	"wms/pkg/utils"
)

type UnitsHandler struct {
	unitsService *unitsservice.UnitsService
}

func NewUnitsHandler(unitsService *unitsservice.UnitsService) *UnitsHandler {
	return &UnitsHandler{
		unitsService: unitsService,
	}
}

func (h *UnitsHandler) GetAll(c *gin.Context) {
	units, err := h.unitsService.GetAll(c)
	if err != nil {
		c.Error(fmt.Errorf("GetAll: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения ед.изм."})
		return
	}

	response := unitsdto.GetAllResponse {
		Units: *units,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UnitsHandler) Create(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Create: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req *unitsdto.CreateRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		c.Error(fmt.Errorf("Create: Некорректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}


	if err := validator.New().Struct(req); err != nil {
		errMsg := utils.ValidationErrors(
			err,
			map[string]string {
				"Name":     "Имя",
				"Sign":     "Обозначение",
				"Code":     "Код",
			},
		)

		c.Error(fmt.Errorf("Create: некорректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	err := h.unitsService.Create(c, req)
	if err != nil {
		c.Error(fmt.Errorf("Create: ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания ед.изм."})
		return
	}

	response := unitsdto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UnitsHandler) Delete(c *gin.Context) {
	var req unitsdto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Delete: ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	err := h.unitsService.Delete(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Delete: неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка удаления ед.изм."})
		return
	}

	response := unitsdto.DeleteResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UnitsHandler) Get(c *gin.Context) {
	var req unitsdto.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Get: неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	user, err := h.unitsService.Get(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Get: неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пользователя"})
		return
	}

	response := unitsdto.GetResponse {
		Name: user.Name,
		Sign: user.Sign,
		Code: user.Code,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UnitsHandler) Set(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Set: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req *unitsdto.SetRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		c.Error(fmt.Errorf("Set: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	if err := validator.New().Struct(req); err != nil {
		errMsg := utils.ValidationErrors(
			err,
			map[string]string{
				"Name":        "Имя",
				"Sign":        "Обозначение",
				"Code":        "Код",
			},
		)

		c.Error(fmt.Errorf("Set: Неккоректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	err := h.unitsService.Set(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := unitsdto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.RouterGroup, unitsService *unitsservice.UnitsService) {
	handler := NewUnitsHandler(unitsService)

	auth := r.Group("/units")
	{
		auth.POST("/getall", handler.GetAll)
		auth.POST("/create", handler.Create)
		auth.POST("/delete", handler.Delete)
		auth.POST("/get", handler.Get)
		auth.POST("/set", handler.Set)
	}
}