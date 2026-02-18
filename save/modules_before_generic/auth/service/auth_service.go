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
)

type AuthService struct {
	userRepo usersrepository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo usersrepository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	user, err := s.userRepo.FindByLogin(ctx, login)
	if err != nil || user == nil {
		return "", fmt.Errorf("Пользователь не существует")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("Неверный пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"login":   user.Login,
		"role_id":    user.RoleID,
		"exp":     time.Now().Add(s.cfg.JWTAccessExpiration).Unix(), 
	})

	tokenString, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("Ошибка генерации токена: %w", err)
	}

	return "Bearer " + tokenString, nil
}

func (s *AuthService) PushToken(c *gin.Context, token string) {
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

func (s *AuthService) ValidateToken(token string) (*jwt.MapClaims, error) {
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