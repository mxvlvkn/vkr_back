package unitsrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/units/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.Unit]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.Unit]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"name", "sign"}, []string{"code"})
	return &repo
}