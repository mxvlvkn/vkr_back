package crudservice

import (
	"context"
	"fmt"

	"wms/internal/config"
	"wms/pkg/crud_module/model"
	"wms/pkg/crud_module/dto"
	"wms/pkg/utils"
	"wms/pkg/crud_module/repository"
)

type CRUDServiceI[
	ModelT crudmodel.Item,
] interface {
	GetAll(ctx context.Context, req cruddto.GetAllRequest) (*[]ModelT, error)
	Create(ctx context.Context, req cruddto.CreateRequest) error
	Delete(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (ModelT, error)
	Set(ctx context.Context, req cruddto.SetRequest) error
	Init(itemRepo crudrepository.ItemRepositoryI[ModelT], cfg *config.Config, tableName string, newItemFunc func()ModelT)
}

type CRUDServiceS[
	ModelT crudmodel.Item,
] struct {
	itemRepo  crudrepository.ItemRepositoryI[ModelT]
	cfg       *config.Config
	tableName string
	newItemFunc func()ModelT
}

func NewGRUDService[
	ModelT crudmodel.Item,
](itemRepo crudrepository.ItemRepositoryI[ModelT], cfg *config.Config, tableName string, newItemFunc func()ModelT) CRUDServiceI[ModelT] {
	return &CRUDServiceS[ModelT]{
		itemRepo: 	  itemRepo,
		cfg:      	  cfg,
		tableName:    tableName,
		newItemFunc:  newItemFunc,
	}
}

func (s *CRUDServiceS[ModelT]) GetAll(ctx context.Context, req cruddto.GetAllRequest) (*[]ModelT, error) {
	items, err := s.itemRepo.GetAll(ctx, req)
	if err != nil || items == nil {
		return nil, fmt.Errorf("Ошибка получения строк \"%v\": %w", s.tableName, err)
	}

	return items, nil
}

func (s *CRUDServiceS[ModelT]) Create(ctx context.Context, req cruddto.CreateRequest) error {
	newItemData := s.newItemFunc()
	utils.FillStructFromStruct(req, newItemData)

	err := s.itemRepo.Create(ctx, newItemData)
	if err != nil {
		return fmt.Errorf("Ошибка создания \"%v\": %w", s.tableName, err)
	}

	return nil
}

func (s *CRUDServiceS[ModelT]) Delete(ctx context.Context, id uint64) error {
	err := s.itemRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("Ошибка удаления \"%v\": %w", s.tableName, err)
	}

	return nil
}

func (s *CRUDServiceS[ModelT]) Get(ctx context.Context, id uint64) (ModelT, error) {
	item, err := s.itemRepo.FindByID(ctx, id)
	if err != nil {
		var nullPtr ModelT
		return nullPtr, fmt.Errorf("Ошибка получения \"%v\": %w", s.tableName, err)
	}

	return item, nil
}

func (s *CRUDServiceS[ModelT]) Set(ctx context.Context, req cruddto.SetRequest) error {
	item := s.newItemFunc()
	updates, err := item.GetUpdateMap(req)
	if err != nil {
		return fmt.Errorf("Ошибка получения updates \"%v\": %w", s.tableName, err)
	}

	err = s.itemRepo.Update(ctx, req.GetID(), updates)
	if err != nil {
		return fmt.Errorf("Ошибка изменения строки \"%v\": %w", s.tableName, err)
	}

	return nil
}

func (s *CRUDServiceS[ModelT]) Init(itemRepo crudrepository.ItemRepositoryI[ModelT], cfg *config.Config, tableName string, newItemFunc func()ModelT) {
	s.itemRepo = itemRepo
	s.cfg = cfg
	s.tableName = tableName
	s.newItemFunc = newItemFunc
}

func (s *CRUDServiceS[ModelT]) GetRepo() crudrepository.ItemRepositoryI[ModelT] {
	return s.itemRepo
}

func (s *CRUDServiceS[ModelT]) GetNewItemFunc() func()ModelT {
	return s.newItemFunc
}