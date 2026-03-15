package marksviewservice

import (
	"wms/internal/config"
	thismodel "wms/internal/modules/marks_view/model"
	thisrepo "wms/internal/modules/marks_view/repository"
	crudservice "wms/pkg/crud_module/service"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.MarkView]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.MarkView]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "MarksView", thismodel.New)
	return &s
}