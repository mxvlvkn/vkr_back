package healthhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler возвращает статус здоровья приложения
// GET /health
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"uptime":    uptime(), // можно добавить функцию расчёта uptime
	})
}

// ReadyHandler — проверка готовности (например, подключение к БД)
// GET /ready
func ReadyHandler(c *gin.Context) {
	// Здесь можно добавить реальную проверку (ping БД, Redis и т.д.)
	// Пока просто возвращаем ok
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// Вспомогательная функция (можно вынести в отдельный файл)
func uptime() string {
	// Если захочешь показывать, сколько работает приложение
	// Нужно хранить startTime в пакете server или глобально
	return "n/a" // заглушка на старте
}