package barcodesviewservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/barcodes_view/model"
	thisrepo "wms/internal/modules/barcodes_view/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.BarcodeView]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.BarcodeView]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "BarcodesView", thismodel.New)
	return &s
}