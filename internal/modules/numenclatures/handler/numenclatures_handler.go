package numenclatureshandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	file, header, _ := c.Request.FormFile(req.ImageFields[0])
	if file != nil {defer file.Close()}

	imageURL, err := h.CRUDService.UploadIMG(c, file, header, "nomenclature")
	if err != nil {
		c.Error(fmt.Errorf("Create: Некорректные данные: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ImageURL = imageURL

	err = h.CRUDService.Create(c, req)
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

func (h *Handler) Delete(c *gin.Context) {
	var req cruddto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Delete: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	numenclature, err := h.CRUDService.Get(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Delete: Ошибка получения удаляемой номенклатуры: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Ошибка получения удаляемой номенклатуры в \"%v\"", h.TableName)})
		return
	}

	err = h.CRUDService.Delete(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Delete: Неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Ошибка удаления \"%v\"", h.TableName)})
		return
	}

	if !strings.Contains(numenclature.ImageURL, "default"){
		err = h.CRUDService.DeleteIMG(numenclature.ImageURL)
		if err != nil {
			c.Error(fmt.Errorf("Delete: Ошибка при удалении фото: %w", err))
		}
	}

	response := cruddto.DeleteResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Set(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Set: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req thisdto.SetRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		c.Error(fmt.Errorf("Set: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	if err := validator.New().Struct(req); err != nil {
		errMsg := utils.ValidationErrors(err)

		c.Error(fmt.Errorf("Set: Неккоректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	numenclature, err := h.CRUDService.Get(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Set: Ошибка получения изменяемой номенклатуры: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Ошибка получения изменяемой номенклатуры в \"%v\"", h.TableName)})
		return
	}

	file, header, _ := c.Request.FormFile(req.ImageFields[0])
	if file != nil {defer file.Close()}

	imageURL, err := h.CRUDService.UploadIMG(c, file, header, "nomenclature")
	if err != nil {
		c.Error(fmt.Errorf("Create: Некорректные данные: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ImageURL = imageURL

	err = h.CRUDService.Set(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !strings.Contains(numenclature.ImageURL, "default"){
		err = h.CRUDService.DeleteIMG(numenclature.ImageURL)
		if err != nil {
			c.Error(fmt.Errorf("Delete: Ошибка при удалении фото: %w", err))
		}
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