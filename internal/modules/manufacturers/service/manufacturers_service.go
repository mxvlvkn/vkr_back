package manufacturersservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/manufacturers/model"
	thisrepo "wms/internal/modules/manufacturers/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.Manufacturer]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.Manufacturer]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Manufacturers", thismodel.New)
	return &s
}