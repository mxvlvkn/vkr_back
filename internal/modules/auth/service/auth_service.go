package authservice

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"wms/internal/config"
	"wms/internal/modules/users/repository"
	"wms/internal/modules/roles/repository"
)

type Service struct {
	userRepo usersrepository.Repository
	roleRepo rolesrepository.Repository
	cfg      *config.Config
}

func NewService(userRepo usersrepository.Repository, roleRepo rolesrepository.Repository, cfg *config.Config) *Service {
	return &Service{
		userRepo: userRepo,
		roleRepo: roleRepo,
		cfg:      cfg,
	}
}

func (s *Service) Login(ctx context.Context, login, password string) (map[string]string, error) {
	user, err := s.userRepo.FindByLogin(ctx, login)
	if err != nil || user == nil {
		return nil, fmt.Errorf("Пользователь не существует")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Неверный пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"login":   user.Login,
		"role_id":    user.RoleID,
		"exp":     time.Now().Add(s.cfg.JWTAccessExpiration).Unix(), 
	})

	role, err := s.roleRepo.FindByID(ctx, user.RoleID)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения роли")
	}

	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("Ошибка генерации токена: %w", err)
	}

	return map[string]string {
		"login": user.Login,
		"role": role.Name,
		"token": "Bearer " + tokenString,
	}, nil
}

func (s *Service) PushToken(c *gin.Context, token string) {
	c.SetCookie(
		"access_token",     							// Имя cookie
		strings.Split(token, " ")[1],              
		int(s.cfg.JWTAccessExpiration / time.Second),          
		"/",                							// Доступ для всего сайта
		"",                 							// Домен (Текущий, если пусто)
		false,               							// Secure (true - только HTTPS)
		true,               							// HttpOnly
	)
}

func (s *Service) ValidateToken(token string) (*jwt.MapClaims, error) {
	tokenRes, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(s.cfg.JWTSecret), nil
	})
	if err != nil || !tokenRes.Valid {
		return nil, errors.New("Недействительный токен")
	}

	claims, ok := tokenRes.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Ошибка получения claims")
	}

	return &claims, nil
}