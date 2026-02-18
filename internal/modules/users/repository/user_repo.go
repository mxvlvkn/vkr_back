package usersrepository

import (
	"context"

	"wms/internal/config"
	thismodel "wms/internal/modules/users/model"
	"wms/pkg/crud_module/repository"

	"gorm.io/gorm"
)

type Repository interface {
	crudrepository.ItemRepositoryI[*thismodel.User]
	FindByLogin(ctx context.Context, login string) (*thismodel.User, error)
}
type repository struct {
	crudrepository.ItemRepositoryS[*thismodel.User]
}

func NewRepository(db *gorm.DB, cfg *config.Config) Repository {
	repo := repository{}
	repo.Init(db, thismodel.New, cfg, []string{}, []string{})
	return &repo
}

func (r *repository) FindByLogin(ctx context.Context, login string) (*thismodel.User, error) {
	var user thismodel.User
	err := r.GetDB().WithContext(ctx).Where("login = ?", login).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}