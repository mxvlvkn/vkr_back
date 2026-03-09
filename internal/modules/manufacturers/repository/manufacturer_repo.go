package manufacturersrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/manufacturers/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.Manufacturer]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.Manufacturer]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"name", "country", "inn", "phone", "email"}, []string{})
	return &repo
}