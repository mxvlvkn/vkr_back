package crudhandler

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"wms/pkg/utils"
	"wms/pkg/crud_module/service"
	"wms/pkg/crud_module/model"
	"wms/pkg/crud_module/dto"
)

type CRUDHandler[
	ModelT crudmodel.Item,
	CreateRequestT cruddto.CreateRequest,
	SetRequestT cruddto.SetRequest,
	GetResponseT cruddto.GetResponse,
] struct {
	CRUDService crudservice.CRUDServiceI[ModelT]
	tableName string
}

func NewCRUDHandler[
	ModelT crudmodel.Item,
	CreateRequestT cruddto.CreateRequest,
	SetRequestT cruddto.SetRequest,
	GetResponseT cruddto.GetResponse,
](
	CRUDService crudservice.CRUDServiceI[ModelT],
	tableName string,
) *CRUDHandler[
	ModelT,
	CreateRequestT,
	SetRequestT,
	GetResponseT,
] {
	return &CRUDHandler[
		ModelT,
		CreateRequestT,
		SetRequestT,
		GetResponseT,
	]{
		CRUDService: CRUDService,
		tableName: tableName,
	}
}

func (h *CRUDHandler[ModelT, CreateRequestT, SetRequestT, GetResponseT]) GetAll(c *gin.Context) {
	var req cruddto.GetAllRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Get: Неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	items, err := h.CRUDService.GetAll(c, req)
	if err != nil {
		c.Error(fmt.Errorf("GetAll: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Ошибка получения \"%v\"", h.tableName),
		})
		return
	}

	response := cruddto.GetAllResponse[ModelT] {
		Items: *items,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CRUDHandler[ModelT, CreateRequestT, SetRequestT, GetResponseT]) Create(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Create: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	fmt.Println("jsonStr")
	fmt.Println(jsonStr)

	var req CreateRequestT
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

	err := h.CRUDService.Create(c, req)
	if err != nil {
		c.Error(fmt.Errorf("Create: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("Ошибка создания \"%v\"", h.tableName)})
		return
	}

	response := cruddto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CRUDHandler[ModelT, CreateRequestT, SetRequestT, GetResponseT]) Delete(c *gin.Context) {
	var req cruddto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Delete: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	err := h.CRUDService.Delete(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Delete: Неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Ошибка удаления \"%v\"", h.tableName)})
		return
	}

	response := cruddto.DeleteResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CRUDHandler[ModelT, CreateRequestT, SetRequestT, GetResponseT]) Get(c *gin.Context) {
	var req cruddto.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Get: Неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	item, err := h.CRUDService.Get(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Get: Неккоректные данные: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("Ошибка получения \"%v\"", h.tableName)})
		return
	}

	response := new(GetResponseT)
	utils.FillStructFromStruct(item, response)

	c.JSON(http.StatusOK, response)
}

func (h *CRUDHandler[ModelT, CreateRequestT, SetRequestT, GetResponseT]) Set(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Set: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req SetRequestT
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

	err := h.CRUDService.Set(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := cruddto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes[
	ModelT crudmodel.Item,
	CreateRequestT cruddto.CreateRequest,
	SetRequestT cruddto.SetRequest,
	GetResponseT cruddto.GetResponse,
](r *gin.RouterGroup, CRUDService crudservice.CRUDServiceI[ModelT], tableName string, route string, useMap map[string]bool) {
	handler := NewCRUDHandler[ModelT, CreateRequestT, SetRequestT, GetResponseT](
		CRUDService,
		tableName,
	)

	api := r.Group("/" + route)
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