package barcodesviewrepository

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/barcodes_view/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.BarcodeView]
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.BarcodeView]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{"code", "numenclature"}, []string{})
	return &repo
}