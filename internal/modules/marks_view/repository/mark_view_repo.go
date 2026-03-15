package marksviewrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/marks_view/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.MarkView]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.MarkView]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"code", "numenclature"}, []string{})
	return &repo
}