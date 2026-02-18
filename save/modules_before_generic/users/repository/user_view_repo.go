package usersrepository

import (
	"context"

	"gorm.io/gorm"
	"wms/internal/modules/users/model"
)

type UserViewRepository interface {
	GetAll(ctx context.Context) (*[]usersmodel.UserView, error)
}

type userViewRepository struct {
	db *gorm.DB
}

func (r *userViewRepository) GetAll(ctx context.Context) (*[]usersmodel.UserView, error) {
	var users []usersmodel.UserView
	err := r.db.WithContext(ctx).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &users, err
}

func NewUserViewRepository(db *gorm.DB) UserViewRepository {
	return &userViewRepository{db: db}
}