package numenclaturesviewrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/numenclatures_view/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.NumenclatureView]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.NumenclatureView]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{}, []string{})
	return &repo
}