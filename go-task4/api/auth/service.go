package auth

import (
	"errors"
	"go-task4/config"
	"go-task4/models"
	"go-task4/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Register(user *models.User) error
	Login(credentials *LoginRequest) (*LoginResponse, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(user *models.User) error {
	// 检查用户名是否已存在
	if _, err := s.repo.FindUserByUsername(user.Username); err == nil {
		return errors.New("username already exists")
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.repo.CreateUser(user)
}

func (s *authService) Login(credentials *LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.FindUserByUsername(credentials.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}
