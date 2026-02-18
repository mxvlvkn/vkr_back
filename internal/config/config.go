package config

import (
	"fmt"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type Config struct {
	Port        		string 			`mapstructure:"PORT"`         
	Environment 		string 			`mapstructure:"ENV"`          
	DBHost      		string			`mapstructure:"DB_HOST"`
	DBPort      		string 			`mapstructure:"DB_PORT"`
	DBUser      		string 			`mapstructure:"DB_USER"`
	DBPassword  		string 			`mapstructure:"DB_PASSWORD"`
	DBName      		string 			`mapstructure:"DB_NAME"`
	DBSSLMode   		string 			`mapstructure:"DB_SSLMODE"`   
	DBMaxOpenConns   	int 			`mapstructure:"DB_MAX_OPEN_CONNS"`   
	DBMaxIdleConns   	int 			`mapstructure:"DB_MAX_IDLE_CONNS"`   
	DBConnMaxLifetime   time.Duration	`mapstructure:"DB_CONN_MAX_LIFETIME"`   
	DBUseMigrate   		bool 			`mapstructure:"DB_USE_MIGRATE"`   
	DBFullLogs   		bool 			`mapstructure:"DB_FULL_LOGS"`   
	JWTSecret   		string 			`mapstructure:"JWT_SECRET"`   
	JWTAccessExpiration time.Duration 	`mapstructure:"JWT_ACCESS_EXPIRATION"`   
	ItemsPageSize		int			 	`mapstructure:"ITEMS_PAGE_SIZE"`   
	AllowedOrigins 		[]string 		`mapstructure:"ALLOWED_ORIGINS"`
	ServerReadTimeout 	int 			`mapstructure:"SERVER_READ_TIMEOUT: 15"`
	ServerWriteTimeout	int 			`mapstructure:"SERVER_WRITE_TIMEOUT: 15"`
	ServerIdleTimeout 	int 			`mapstructure:"SERVER_IDLE_TIMEOUT: 60"`
}

// Загружаем конфигурацию из файла
func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")           
	v.SetConfigType("yaml")             
	v.AddConfigPath("configs")        

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Ошибка чтения конфига: %w", err)
	}

	var cfg Config
	decoderConfig := func(dc *mapstructure.DecoderConfig) {dc.DecodeHook = mapstructure.StringToTimeDurationHookFunc()}
	if err := v.Unmarshal(&cfg, decoderConfig); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга конфигурации: %w", err)
	}

	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("Переменные БД не заданы: DB_HOST, DB_USER, DB_NAME")
	}

	return &cfg, nil
}