package unitsservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/units/model"
	thisrepo "wms/internal/modules/units/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.Unit]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.Unit]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Units", thismodel.New)
	return &s
}