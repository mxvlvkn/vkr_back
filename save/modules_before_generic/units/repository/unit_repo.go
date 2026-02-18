package unitsrepository

import (
	"context"

	"wms/internal/modules/units/model"


	"gorm.io/gorm"
)

type UnitRepository interface {
	GetAll(ctx context.Context) (*[]unitsmodel.Unit, error)
	Create(ctx context.Context, newUnitData *unitsmodel.Unit) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*unitsmodel.Unit, error)
	Update(ctx context.Context, unitID uint64, updates map[string]any) error
}

type unitRepository struct {
	db *gorm.DB
}

func (r *unitRepository) GetAll(ctx context.Context) (*[]unitsmodel.Unit, error) {
	var units []unitsmodel.Unit
	err := r.db.WithContext(ctx).Find(&units).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &units, err
}

func (r *unitRepository) Create(ctx context.Context, newUnitData *unitsmodel.Unit) error {
	return r.db.WithContext(ctx).Create(newUnitData).Error
}

func (r *unitRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&unitsmodel.Unit{}, id).Error
}

func (r *unitRepository) FindByID(ctx context.Context, id uint64) (*unitsmodel.Unit, error) {
	var unit unitsmodel.Unit
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&unit).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &unit, err
}

func (r *unitRepository) Update(ctx context.Context, unitID uint64, updates map[string]any) error {
    return r.db.WithContext(ctx).
        Model(&unitsmodel.Unit{ID: unitID}). 
        Updates(updates).Error 
}

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepository{db: db}
}