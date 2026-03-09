package app

import (
	"context"
	"fmt"


	"gorm.io/gorm"


	"wms/internal/config"
	"wms/internal/database"
	"wms/internal/middleware"
	"wms/internal/server"

	"wms/internal/modules/auth/service"
	"wms/internal/modules/auth/handler"

	"wms/internal/modules/roles/repository"
	"wms/internal/modules/roles/service"
	"wms/internal/modules/roles/handler"

	"wms/internal/modules/units/repository"
	"wms/internal/modules/units/service"
	"wms/internal/modules/units/handler"

    "wms/internal/modules/users/repository"
	"wms/internal/modules/users/service"
	"wms/internal/modules/users/handler"

    "wms/internal/modules/users_view/repository"
	"wms/internal/modules/users_view/service"
	"wms/internal/modules/users_view/handler"

    "wms/internal/modules/manufacturers/repository"
	"wms/internal/modules/manufacturers/service"
	"wms/internal/modules/manufacturers/handler"

    "wms/internal/modules/numenclatures/repository"
	"wms/internal/modules/numenclatures/service"
	"wms/internal/modules/numenclatures/handler"

    "wms/internal/modules/numenclatures_view/repository"
	"wms/internal/modules/numenclatures_view/service"
	"wms/internal/modules/numenclatures_view/handler"
)


type App struct {
	cfg    *config.Config
	server *server.Server
	db     *gorm.DB
}

// Создание и инициализация приложения
func New() (*App, error) {
    cfg, err := config.Load()
    if err != nil {
        return nil, fmt.Errorf("Не удалось загрузить конфигурацию: %w", err)
    }

    db, err := database.NewConnection(cfg)
    if err != nil {
        return nil, err
    }

    // Запуск миграций
    if cfg.DBUseMigrate {
        if err := database.MigrateDB(cfg); err != nil {
            return nil, fmt.Errorf("Миграции не применились: %w", err)
        }
    }

    // Создание репозиториев
    userRepo := usersrepository.NewRepository(db, cfg)
    userViewRepo := usersviewrepository.NewRepository(db, cfg)
    roleRepo := rolesrepository.NewRepository(db, cfg)
    unitRepo := unitsrepository.NewRepository(db, cfg)
    manufacturerRepo := manufacturersrepository.NewRepository(db, cfg)
    numenclatureRepo := numenclaturesrepository.NewRepository(db, cfg)
    numenclatureViewRepo := numenclaturesviewrepository.NewRepository(db, cfg)

    // Создание сервисов
    authService := authservice.NewService(userRepo, roleRepo, cfg)
    usersService := usersservice.NewService(userRepo, cfg)
    usersViewService := usersviewservice.NewService(userViewRepo, cfg)
    rolesService := rolesservice.NewService(roleRepo, cfg)
    unitsService := unitsservice.NewService(unitRepo, cfg)
    manufacturersService := manufacturersservice.NewService(manufacturerRepo, cfg)
    numenclaturesService := numenclaturesservice.NewService(numenclatureRepo, cfg)
    numenclaturesViewService := numenclaturesviewservice.NewService(numenclatureViewRepo, cfg)

    // Создание сервера
    srv := server.New(cfg, db)

    // Создание роутов
    api := srv.Router().Group("/api")
    authhandler.RegisterRoutes(api, authService)

	api.Use(middleware.AuthMiddleware(authService))

    usershandler.RegisterRoutes(api, usersService)
    usersviewhandler.RegisterRoutes(api, usersViewService)
    roleshandler.RegisterRoutes(api, rolesService)
    unitshandler.RegisterRoutes(api, unitsService)
    manufacturershandler.RegisterRoutes(api, manufacturersService)
    numenclatureshandler.RegisterRoutes(api, numenclaturesService)
    numenclaturesviewhandler.RegisterRoutes(api, numenclaturesViewService)

    return &App{
        cfg:    cfg,
        server: srv,
        db:     db,
    }, nil
}

// Запускаем HTTP-сервер
func (a *App) Run() error {
	return a.server.Run(a.cfg.Port)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}