package marksrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/marks/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.Mark]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.Mark]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"code"}, []string{})
	return &repo
}