package barcodesservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/barcodes/model"
	thisrepo "wms/internal/modules/barcodes/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.Barcode]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.Barcode]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Barcodes", thismodel.New)
	return &s
}