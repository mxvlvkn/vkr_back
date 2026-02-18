package rolesrepository

import (
	"context"

	"gorm.io/gorm"
	"wms/internal/modules/roles/model"
)

type RoleRepository interface {
	GetAll(ctx context.Context) (*[]rolesmodel.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r *roleRepository) GetAll(ctx context.Context) (*[]rolesmodel.Role, error) {
	var roles []rolesmodel.Role
	err := r.db.WithContext(ctx).Find(&roles).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &roles, err
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}