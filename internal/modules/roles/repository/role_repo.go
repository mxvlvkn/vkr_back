package rolesrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/roles/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.Role]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.Role]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{}, []string{})
	return &repo
}