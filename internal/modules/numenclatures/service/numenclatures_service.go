package numenclaturesservice

import (
	"context"
	"fmt"
	"wms/internal/config"
	thismodel "wms/internal/modules/numenclatures/model"
	thisrepo "wms/internal/modules/numenclatures/repository"
	cruddto "wms/pkg/crud_module/dto"
	crudservice "wms/pkg/crud_module/service"
	"wms/pkg/utils"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.Numenclature]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.Numenclature]
}

func (s *ServiceS) Create(ctx context.Context, req cruddto.CreateRequest) error {
	newItemData := s.NewItemFunc()
	utils.FillStructFromStruct(req, newItemData)

	//!!!!!!!!!!!

	err := s.ItemRepo.Create(ctx, newItemData)
	if err != nil {
		return fmt.Errorf("Ошибка создания \"%v\": %w", s.TableName, err)
	}

	return nil
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Numenclatures", thismodel.New)
	return &s
}