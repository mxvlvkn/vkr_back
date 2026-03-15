package crudrepository

import (
	"context"
	"strings"
	"fmt"

	"wms/internal/config"
	"wms/pkg/crud_module/dto"
	"wms/pkg/crud_module/model"
	"wms/pkg/utils"

	"gorm.io/gorm"
)

type ItemRepositoryI[
	ModelT crudmodel.Item,
] interface {
	GetDB() *gorm.DB
	GetAll(ctx context.Context, req cruddto.GetAllRequest) (*[]ModelT, error)
	Create(ctx context.Context, newItemData ModelT) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (ModelT, error)
	Update(ctx context.Context, itemID uint64, updates map[string]any) error
}

type ItemRepositoryS[
	ModelT crudmodel.Item,
] struct {
	db *gorm.DB
	newItemFunc func()ModelT
	cfg *config.Config
	searchStrFields []string
	searchIntFields []string
}

func (r *ItemRepositoryS[ModelT]) GetAll(ctx context.Context, req cruddto.GetAllRequest) (*[]ModelT, error) {
	var items []ModelT
	pageNum := req.PageNum

	db := r.db.WithContext(ctx)
	if (pageNum != 0) {
		pageSize := r.cfg.ItemsPageSize
		offset := (pageNum - 1) * pageSize

		db = db.
			Limit(pageSize).
			Offset(offset)
	}

	if req.WhereID != -1 && req.WhereField != "" {
		db.Where(fmt.Sprintf("%s = ?", req.WhereField), req.WhereID)
	}

	if req.Search != "" {
		var conditions []string
		var args []any

		for _, col := range r.searchStrFields {
			conditions = append(conditions, col+" ILIKE ?")
			args = append(args, "%"+req.Search+"%")
		}

		for _, col := range r.searchIntFields {
			conditions = append(conditions, "CAST("+col+" AS TEXT) LIKE ?")
			args = append(args, "%"+req.Search+"%")
		}

		db = db.Where(strings.Join(conditions, " OR "), args...)
	}

	if req.FilterMethod != "default" && req.FilterField != "" {
		order := utils.CamelCaseToSnake(req.FilterField) + " " + req.FilterMethod
    	db = db.Order(order)
	}

	err := db.Find(&items).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &items, err
}

func (r *ItemRepositoryS[ModelT]) Create(ctx context.Context, newItemData ModelT) error {
	return r.db.WithContext(ctx).Create(newItemData).Error
}

func (r *ItemRepositoryS[ModelT]) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(r.newItemFunc(), id).Error
}

func (r *ItemRepositoryS[ModelT]) FindByID(ctx context.Context, id uint64) (ModelT, error) {
	item := r.newItemFunc()
	err := r.db.WithContext(ctx).Where("id = ?", id).First(item).Error
	if err == gorm.ErrRecordNotFound {
        var nullPtr ModelT
		return nullPtr, nil
	}
	if err != nil {
        var nullPtr ModelT
        return nullPtr, err
    }
	return item, err
}

func (r *ItemRepositoryS[ModelT]) Update(ctx context.Context, itemID uint64, updates map[string]any) error {
	item := r.newItemFunc()
	utils.FillStructFromStruct(
		struct{ID uint64}{ID: itemID},
		item,
	)
    return r.db.WithContext(ctx).
        Model(item). 
        Updates(updates).Error 
}

func NewCRUDRepository[ModelT crudmodel.Item](
	db *gorm.DB, 
	newItemFunc func()ModelT, 
	cfg *config.Config,
	searchStrFields []string,
	searchIntFields []string,
) ItemRepositoryI[ModelT] {
	return &ItemRepositoryS[ModelT]{
		db: db, 
		newItemFunc: newItemFunc, 
		cfg: cfg,
		searchStrFields: searchStrFields,
		searchIntFields: searchIntFields,
	}
}

func (r *ItemRepositoryS[ModelT]) GetDB() *gorm.DB {
	return r.db
}

func (r *ItemRepositoryS[ModelT]) Init(
	db *gorm.DB, 
	newItemFunc func()ModelT, 
	cfg *config.Config,
	searchStrFields []string,
	searchIntFields []string, 
) {
	r.db = db
	r.newItemFunc = newItemFunc
	r.cfg = cfg
	r.searchStrFields = searchStrFields
	r.searchIntFields = searchIntFields
}

