package rolesservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/roles/model"
	thisrepo "wms/internal/modules/roles/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.Role]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.Role]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Roles", thismodel.New)
	return &s
}
