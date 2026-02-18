package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"wms/internal/app"
)

func main() {
	// Создание и инициализация приложения
	application, err := app.New()
	if err != nil {
		log.Fatalf("Не удалось инициализировать приложение: %v", err)
	}

	go func() {
		if err := application.Run(); err != nil {
			log.Fatalf("Ошибка при работе сервера: %v", err)
		}
	}()

	// Ожидание сигнала завершения (Ctrl+C, SIGTERM от Docker/K8s)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при остановке сервера: %v", err)
	} else {
		log.Println("Приложение успешно завершило работу")
	}
}