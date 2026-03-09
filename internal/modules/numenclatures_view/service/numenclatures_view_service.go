package numenclaturesviewservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/numenclatures_view/model"
	thisrepo "wms/internal/modules/numenclatures_view/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.NumenclatureView]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.NumenclatureView]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "NumenclaturesView", thismodel.New)
	return &s
}