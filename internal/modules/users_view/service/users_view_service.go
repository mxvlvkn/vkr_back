package usersviewservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/users_view/model"
	thisrepo "wms/internal/modules/users_view/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.UserView]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.UserView]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "UsesrView", thismodel.New)
	return &s
}