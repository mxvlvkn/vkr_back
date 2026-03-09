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
	Init(ItemRepo crudrepository.ItemRepositoryI[ModelT], cfg *config.Config, TableName string, newItemFunc func()ModelT)
}

type CRUDServiceS[
	ModelT crudmodel.Item,
] struct {
	ItemRepo  crudrepository.ItemRepositoryI[ModelT]
	cfg       *config.Config
	TableName string
	NewItemFunc func()ModelT
}

func NewGRUDService[
	ModelT crudmodel.Item,
](ItemRepo crudrepository.ItemRepositoryI[ModelT], cfg *config.Config, TableName string, newItemFunc func()ModelT) CRUDServiceI[ModelT] {
	return &CRUDServiceS[ModelT]{
		ItemRepo: 	  ItemRepo,
		cfg:      	  cfg,
		TableName:    TableName,
		NewItemFunc:  newItemFunc,
	}
}

func (s *CRUDServiceS[ModelT]) GetAll(ctx context.Context, req cruddto.GetAllRequest) (*[]ModelT, error) {
	items, err := s.ItemRepo.GetAll(ctx, req)
	if err != nil || items == nil {
		return nil, fmt.Errorf("Ошибка получения строк \"%v\": %w", s.TableName, err)
	}

	return items, nil
}

func (s *CRUDServiceS[ModelT]) Create(ctx context.Context, req cruddto.CreateRequest) error {
	newItemData := s.NewItemFunc()
	utils.FillStructFromStruct(req, newItemData)

	err := s.ItemRepo.Create(ctx, newItemData)
	if err != nil {
		return fmt.Errorf("Ошибка создания \"%v\": %w", s.TableName, err)
	}

	return nil
}

func (s *CRUDServiceS[ModelT]) Delete(ctx context.Context, id uint64) error {
	err := s.ItemRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("Ошибка удаления \"%v\": %w", s.TableName, err)
	}

	return nil
}

func (s *CRUDServiceS[ModelT]) Get(ctx context.Context, id uint64) (ModelT, error) {
	item, err := s.ItemRepo.FindByID(ctx, id)
	if err != nil {
		var nullPtr ModelT
		return nullPtr, fmt.Errorf("Ошибка получения \"%v\": %w", s.TableName, err)
	}

	return item, nil
}

func (s *CRUDServiceS[ModelT]) Set(ctx context.Context, req cruddto.SetRequest) error {
	item := s.NewItemFunc()
	updates, err := item.GetUpdateMap(req)
	if err != nil {
		return fmt.Errorf("Ошибка получения updates \"%v\": %w", s.TableName, err)
	}

	err = s.ItemRepo.Update(ctx, req.GetID(), updates)
	if err != nil {
		return fmt.Errorf("Ошибка изменения строки \"%v\": %w", s.TableName, err)
	}

	return nil
}

func (s *CRUDServiceS[ModelT]) Init(ItemRepo crudrepository.ItemRepositoryI[ModelT], cfg *config.Config, TableName string, newItemFunc func()ModelT) {
	s.ItemRepo = ItemRepo
	s.cfg = cfg
	s.TableName = TableName
	s.NewItemFunc = newItemFunc
}

func (s *CRUDServiceS[ModelT]) GetRepo() crudrepository.ItemRepositoryI[ModelT] {
	return s.ItemRepo
}

func (s *CRUDServiceS[ModelT]) GetNewItemFunc() func()ModelT {
	return s.NewItemFunc
}