package numenclatureshandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	thisdto "wms/internal/modules/numenclatures/dto"
	thismodel "wms/internal/modules/numenclatures/model"
	thisservice "wms/internal/modules/numenclatures/service"
	cruddto "wms/pkg/crud_module/dto"
	crudhandler "wms/pkg/crud_module/handler"
	"wms/pkg/utils"
)

type Handler struct {
	crudhandler.CRUDHandler[
		*thismodel.Numenclature,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	]
} 

func NewHandler(service thisservice.ServiceI) *Handler {
	base := crudhandler.NewCRUDHandler[
		*thismodel.Numenclature,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Numenclatures")

	return &Handler{
		*base,
	}
}

func (h *Handler) Create(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Create: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req thisdto.CreateRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		c.Error(fmt.Errorf("Create: Некорректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}


	if err := validator.New().Struct(req); err != nil {
		errMsg := utils.ValidationErrors(err)

		c.Error(fmt.Errorf("Create: Некорректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	file, header, _ := c.Request.FormFile("image")
	req.Image = file
	req.ImageHeader = header

	err := h.CRUDService.Create(c, req)
	if err != nil {
		c.Error(fmt.Errorf("Create: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка создания \"%v\"", h.TableName)})
		return
	}

	response := cruddto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	handler := NewHandler(service)
	useMap := map[string]bool {
		"getall": 	false,
		"create": 	true,
		"delete": 	true,
		"get": 		true,
		"set": 		true,
	}

	api := r.Group("/numenclatures")
	if v, ok := useMap["getall"]; ok && v {
		api.POST("/getall", handler.GetAll)
	}
	if v, ok := useMap["create"]; ok && v {
		api.POST("/create", handler.Create)
	}
	if v, ok := useMap["delete"]; ok && v {
		api.POST("/delete", handler.Delete)
	}
	if v, ok := useMap["get"]; ok && v {
		api.POST("/get", handler.Get)
	}
	if v, ok := useMap["set"]; ok && v {
		api.POST("/set", handler.Set)
	}
}