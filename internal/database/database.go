package database

import (
	"fmt"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "gorm.io/gorm/logger"

    "wms/internal/config"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        cfg.DBHost, 
        cfg.DBUser, 
        cfg.DBPassword, 
        cfg.DBName, 
        cfg.DBPort, 
        cfg.DBSSLMode,
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config {
        Logger: logger.Default.LogMode(getLoggerMode(cfg.DBFullLogs)),
    })
    if err != nil {
        return nil, fmt.Errorf("Не удалось подключиться к БД (приложение): %w", err)
    }

	// Настройка пула соединений
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
    sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
    sqlDB.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	return db, nil
}

func getLoggerMode(useFullLogs bool) logger.LogLevel {
    if useFullLogs {
        return logger.Info
    } else {
        return logger.Silent
    }
}