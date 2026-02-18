package usersservice

import (
	"context"
	"fmt"

	"wms/internal/config"
	"wms/internal/modules/users/dto"
	"wms/internal/modules/users/model"
	"wms/internal/modules/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	userViewRepo usersrepository.UserViewRepository
	userRepo usersrepository.UserRepository
	cfg      *config.Config
}

func NewUsersService(userRepo usersrepository.UserRepository, userViewRepo usersrepository.UserViewRepository, cfg *config.Config) *UsersService {
	return &UsersService{
		userRepo: userRepo,
		userViewRepo: userViewRepo,
		cfg:      cfg,
	}
}

func (s *UsersService) GetAll(ctx context.Context) (*[]usersmodel.UserView, error) {
	users, err := s.userViewRepo.GetAll(ctx)
	if err != nil || users == nil {
		return nil, fmt.Errorf("Ошибка получения пользователей")
	}

	return users, nil
}

func (s *UsersService) Create(ctx context.Context, req *usersdto.CreateRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("Ошибка хэширования пароля")
    }

	newUserData := &usersmodel.User {
		Login: req.Login,
		Name: req.Name,
		Surname: req.Surname,
		Patronymic: req.Patronymic,
		RoleID: req.RoleID,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepo.Create(ctx, newUserData)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Ошибка создания пользователя")
	}

	return nil
}

func (s *UsersService) Delete(ctx context.Context, id uint64) error {
	err := s.userRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("Ошибка удаления пользователя")
	}

	return nil
}

func (s *UsersService) Get(ctx context.Context, id uint64) (*usersmodel.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения пользователя")
	}

	return user, nil
}

func (s *UsersService) Set(ctx context.Context, req *usersdto.SetRequest) error {
	updates := map[string]any {
		"login": req.Login,
		"name": req.Name,
		"surname": req.Surname,
		"patronymic": req.Patronymic,
		"role_id": req.RoleID,
	}

	if req.IsSetPassword {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("Ошибка хэширования пароля")
		}

		updates["password_hash"] = hashedPassword
	}

	err := s.userRepo.Update(ctx, req.ID, updates)
	if err != nil {
		return fmt.Errorf("Ошибка создания пользователя")
	}

	return nil
}