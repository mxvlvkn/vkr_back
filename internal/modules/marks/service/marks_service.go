package marksservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/marks/model"
	thisrepo "wms/internal/modules/marks/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.Mark]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.Mark]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Marks", thismodel.New)
	return &s
}