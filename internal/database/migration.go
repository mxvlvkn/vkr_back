package database

import (
	"fmt"
	"log"
	"wms/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDB(cfg *config.Config) error {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        cfg.DBUser, 
        cfg.DBPassword, 
        cfg.DBHost, 
        cfg.DBPort, 
        cfg.DBName, 
        cfg.DBSSLMode,
    )

	m, err := migrate.New(
		"file://migrations",  
		dsn,
	)
	if err != nil {
		return fmt.Errorf("Не удалось создать миграцию: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Миграции: изменений нет (БД актуальна)")
			return nil
		}
		return fmt.Errorf("ошибка применения миграций: %w", err)
	}

	log.Println("Миграции успешно применены")
	return nil
}