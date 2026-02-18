package usersrepository

import (
	"context"

	"wms/internal/modules/users/model"


	"gorm.io/gorm"
)

type UserRepository interface {
	FindByLogin(ctx context.Context, login string) (*usersmodel.User, error)
	Create(ctx context.Context, newUserData *usersmodel.User) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*usersmodel.User, error)
	Update(ctx context.Context, userID uint64, updates map[string]any) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindByLogin(ctx context.Context, login string) (*usersmodel.User, error) {
	var user usersmodel.User
	err := r.db.WithContext(ctx).Where("login = ?", login).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *userRepository) Create(ctx context.Context, newUserData *usersmodel.User) error {
	return r.db.WithContext(ctx).Create(newUserData).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&usersmodel.User{}, id).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint64) (*usersmodel.User, error) {
	var user usersmodel.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *userRepository) Update(ctx context.Context, userID uint64, updates map[string]any) error {
    return r.db.WithContext(ctx).
        Model(&usersmodel.User{ID: userID}). 
        Updates(updates).Error 
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}