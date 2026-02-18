package rolesservice

import (
	"context"
	"errors"

	"wms/internal/config"
	"wms/internal/modules/roles/model"
	"wms/internal/modules/roles/repository"
)

type RolesService struct {
	roleRepo rolesrepository.RoleRepository
	cfg      *config.Config
}

func NewRolesService(roleRepo rolesrepository.RoleRepository, cfg *config.Config) *RolesService {
	return &RolesService{
		roleRepo: roleRepo,
		cfg:      cfg,
	}
}

func (s *RolesService) GetAll(ctx context.Context) (*[]rolesmodel.Role, error) {
	users, err := s.roleRepo.GetAll(ctx)
	if err != nil || users == nil {
		return nil, errors.New("Ошибка получения ролей")
	}

	return users, nil
}