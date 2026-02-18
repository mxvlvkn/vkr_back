package usersservice

import (
	"context"
	"fmt"

	"wms/pkg/utils"
	"wms/internal/config"
	thismodel "wms/internal/modules/users/model"
	thisrepo "wms/internal/modules/users/repository"
	cruddto "wms/pkg/crud_module/dto"
	crudservice "wms/pkg/crud_module/service"

	"golang.org/x/crypto/bcrypt"
)

type ServiceI interface {
	crudservice.CRUDServiceI[*thismodel.User]
}

type ServiceS struct {
	crudservice.CRUDServiceS[*thismodel.User]
}

func NewService(repo thisrepo.Repository, cfg *config.Config) ServiceI {
	s := ServiceS{}
	s.Init(repo, cfg, "Users", thismodel.New)
	return &s
}

func (s *ServiceS) Create(ctx context.Context, req cruddto.CreateRequest) error {
	newItemData := s.GetNewItemFunc()()
	utils.FillStructFromStruct(req, newItemData)

	passwordField, err := utils.GetStructFieldByString(req, "Password")
	if err != nil {
		return fmt.Errorf("Ошибка получения passwordField")
	}
	password, ok := passwordField.(string)
	if !ok {
		return fmt.Errorf("Ошибка получения пароля")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Ошибка хэширования пароля")
	}

	newItemData.PasswordHash = string(hashedPassword)

	err = s.GetRepo().Create(ctx, newItemData)
	if err != nil {
		return fmt.Errorf("Ошибка создания \"Users\": %w", err)
	}

	return nil
}

func (s *ServiceS) Set(ctx context.Context, req cruddto.SetRequest) error {
	item := s.GetNewItemFunc()()
	updates, err := item.GetUpdateMap(req)
	if err != nil {
		return fmt.Errorf("Ошибка получения updates \"Users\": %w", err)
	}

	isSetPasswordField, err := utils.GetStructFieldByString(req, "IsSetPassword")
	if err != nil {
		return fmt.Errorf("Ошибка получения isSetPasswordField")
	}
	isSetPassword, ok := isSetPasswordField.(bool)
	if !ok {
		return fmt.Errorf("Ошибка получения isSetPassword")
	}
	
	passwordField, err := utils.GetStructFieldByString(req, "Password")
	if err != nil {
		return fmt.Errorf("Ошибка получения passwordField")
	}
	password, ok := passwordField.(string)
	if !ok {
		return fmt.Errorf("Ошибка получения пароля")
	}

	if bool(isSetPassword) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("Ошибка хэширования пароля")
		}

		updates["password_hash"] = hashedPassword
	}

	err = s.GetRepo().Update(ctx, req.GetID(), updates)
	if err != nil {
		return fmt.Errorf("Ошибка изменения строки \"Users\": %w", err)
	}
	
	return nil
}