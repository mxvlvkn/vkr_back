package usersviewrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/users_view/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.UserView]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.UserView]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"login", "name", "surname", "patronymic", "role"}, []string{})
	return &repo
}