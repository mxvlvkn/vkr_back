package numenclaturesrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/numenclatures/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.Numenclature]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.Numenclature]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"name", "article"}, []string{})
	return &repo
}