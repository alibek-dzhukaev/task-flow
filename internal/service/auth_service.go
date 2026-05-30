package service

import (
	"errors"
	"time"

	"github.com/alibek-dzhukaev/task-flow/internal/model"
	"github.com/alibek-dzhukaev/task-flow/internal/repository"
	"github.com/alibek-dzhukaev/task-flow/internal/util"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(email, password, firstName, lastName string) (*model.User, error)
	Login(email, password string) (accessToken string, err error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
	tokenTTL  time.Duration
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string, tokenTTL time.Duration) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		tokenTTL:  tokenTTL,
	}
}

func (s *authService) Register(email, password, firstName, lastName string) (*model.User, error) {
	existing, err := s.userRepo.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hash, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:     email,
		Password:  hash,
		FirstName: firstName,
		LastName:  lastName,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(email, password string) (accessToken string, err error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !util.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return util.GenerateAccessToken(user.ID, user.Email, user.Role, s.jwtSecret, s.tokenTTL)
}
