package usershandler

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"wms/internal/modules/users/dto"
	"wms/internal/modules/users/service"
	"wms/pkg/utils"
)

type UsersHandler struct {
	usersService *usersservice.UsersService
}

func NewUsersHandler(usersService *usersservice.UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

func (h *UsersHandler) GetAll(c *gin.Context) {
	users, err := h.usersService.GetAll(c)
	if err != nil {
		c.Error(fmt.Errorf("GetAll: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := usersdto.GetAllResponse {
		Users: *users,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Create(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Create: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req *usersdto.CreateRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		c.Error(fmt.Errorf("Create: Некорректные данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := validator.New().Struct(req); err != nil {
		errMsg := utils.ValidationErrors(
			err,
			map[string]string{
				"Login":          "Логин",
				"Name":           "Имя",
				"Surname":        "Фамилия",
				"Patronymic":     "Отчество",
				"RoleID":         "Должность",
				"Password":       "Пароль",
				"RepeatPassword": "Повторение пароля",
			},
		)

		c.Error(fmt.Errorf("Create: Некорректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	err := h.usersService.Create(c, req)
	if err != nil {
		c.Error(fmt.Errorf("Create: ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := usersdto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Delete(c *gin.Context) {
	var req usersdto.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := utils.ValidationErrors(
			err,
			map[string]string{
				"Login": "Логин",
			},
		)

		c.Error(fmt.Errorf("Delete: Неккоректные данные: %v", errMsg))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	err := h.usersService.Delete(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Delete: ошибка обработки запроса: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пользователя"})
		return
	}

	response := usersdto.DeleteResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Get(c *gin.Context) {
	var req usersdto.GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(fmt.Errorf("Get: Неккоректные данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неккоректные данные"})
		return
	}

	user, err := h.usersService.Get(c.Request.Context(), req.ID)
	if err != nil {
		c.Error(fmt.Errorf("Get: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := usersdto.GetResponse {
		Login: user.Login,
		Name: user.Name,
		Surname: user.Surname,
		Patronymic: user.Patronymic,
		RoleID: user.RoleID,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) Set(c *gin.Context) {
	jsonStr := c.PostForm("data")
	if jsonStr == "" {
		c.Error(fmt.Errorf("Set: Пустые данные"))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Пустые данные"})
		return
	}

	var req *usersdto.SetRequest
	if err := json.Unmarshal([]byte(jsonStr), &req); err != nil {
		c.Error(fmt.Errorf("Set: Некорректные данные"))
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "Некорректные данные"},
		)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		errMsg := utils.ValidationErrors(
			err,
			map[string]string{
				"Login":          "Логин",
				"Name":           "Имя",
				"Surname":        "Фамилия",
				"Patronymic":     "Отчество",
				"RoleID":         "Должность",
				"Password":       "Пароль",
				"RepeatPassword": "Повторение пароля",
			},
		)

		c.Error(fmt.Errorf("Get: Неккоректные данные: %v", errMsg))
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error": errMsg,
			},
		)
		return
	}

	err := h.usersService.Set(c, req)
	if err != nil {
		c.Error(fmt.Errorf("Set: Ошибка обработки запроса: %w", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка изменения пользователя"})
		return
	}

	response := usersdto.CreateResponse {
		Status: true,
	}

	c.JSON(http.StatusOK, response)
}

func RegisterRoutes(r *gin.RouterGroup, usersService *usersservice.UsersService) {
	handler := NewUsersHandler(usersService)

	auth := r.Group("/users")
	{
		auth.POST("/getall", handler.GetAll)
		auth.POST("/create", handler.Create)
		auth.POST("/delete", handler.Delete)
		auth.POST("/get", handler.Get)
		auth.POST("/set", handler.Set)
	}
}