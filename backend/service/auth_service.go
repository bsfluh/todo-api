package service

import (
	"fmt"
	"time"
	"todo-Api/model"
	"todo-Api/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(name, email, password string) error
	Login(email, password string) (*model.User, error)
}
type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		userRepo: repo,
	}
}
func (s *authService) Register(name, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := model.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}
	return s.userRepo.Create(&user)
}

func (s *authService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	return user, nil
}
