package unitsservice

import (
	"context"
	"fmt"

	"wms/internal/config"
	// "wms/internal/modules/units/dto"
	unitsdto "wms/internal/modules/units/dto"
	"wms/internal/modules/units/model"
	"wms/internal/modules/units/repository"
)

type UnitsService struct {
	unitRepo unitsrepository.UnitRepository
	cfg      *config.Config
}

func NewUnitsService(unitRepo unitsrepository.UnitRepository, cfg *config.Config) *UnitsService {
	return &UnitsService{
		unitRepo: unitRepo,
		cfg:      cfg,
	}
}

func (s *UnitsService) GetAll(ctx context.Context) (*[]unitsmodel.Unit, error) {
	users, err := s.unitRepo.GetAll(ctx)
	if err != nil || users == nil {
		return nil, fmt.Errorf("Ошибка получения ед.изм.")
	}

	return users, nil
}

func (s *UnitsService) Create(ctx context.Context, req *unitsdto.CreateRequest) error {
	newUnitData := &unitsmodel.Unit {
		Name: req.Name,
		Sign: req.Sign,
		Code: req.Code,
	}

	err := s.unitRepo.Create(ctx, newUnitData)
	if err != nil {
		return fmt.Errorf("Ошибка создания ед.изм.: %w", err)
	}

	return nil
}

func (s *UnitsService) Delete(ctx context.Context, id uint64) error {
	err := s.unitRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("Ошибка удаления ед.изм.")
	}

	return nil
}

func (s *UnitsService) Get(ctx context.Context, id uint64) (*unitsmodel.Unit, error) {
	unit, err := s.unitRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения ед.изм.")
	}

	return unit, nil
}

func (s *UnitsService) Set(ctx context.Context, req *unitsdto.SetRequest) error {
	updates := map[string]any {
		"name": req.Name,
		"sign": req.Sign,
		"code": req.Code,
	}

	err := s.unitRepo.Update(ctx, req.ID, updates)
	if err != nil {
		return fmt.Errorf("Ошибка создания ед.изм.")
	}

	return nil
}