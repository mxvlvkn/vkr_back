package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"wms/internal/config"
	"wms/internal/middleware"
	"wms/internal/server/handler"
)

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
}

func New(cfg *config.Config, db *gorm.DB) *Server {
	if cfg.Environment == "production" || cfg.Environment == "staging" {
		gin.SetMode(gin.ReleaseMode) 
	} else {
		gin.SetMode(gin.DebugMode) 
	}

	router := gin.New()
	// Восстановливаем после паники
	router.Use(gin.Recovery()) 
	// 
	router.Use(middleware.ErrorLogger())
	router.Use(middleware.CORS(cfg))

	// Глобальные обработчики
	//!router.GET("/", handler.RootHandler)
	router.GET("/health", healthhandler.HealthHandler)
	router.GET("/ready", healthhandler.ReadyHandler)
	router.Static("/uploads", "./uploads")

	
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ServerReadTimeout) * time.Second,  
		WriteTimeout: time.Duration(cfg.ServerWriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.ServerIdleTimeout) * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		router:     router,
	}
}

// Запуск сервера
func (s *Server) Run(port string) error {
	s.httpServer.Addr = ":" + port

	fmt.Printf("Сервер запущен на http://localhost%s\n", s.httpServer.Addr)
	fmt.Printf("Health check: http://localhost%s/health\n", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// Router возвращает gin.Engine для регистрации роутов из app.go
func (s *Server) Router() *gin.Engine {
    return s.router
}