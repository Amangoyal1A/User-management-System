package service

import (
	"errors"
	"fmt"
	"user-management/models"
	"user-management/repository"
	"user-management/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (string, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) (*models.User, error) {
	hashedPw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPw)
	res, err := s.repo.CreateUser(user)
	return res, err
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.ID)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id uint, updatedUser *models.User) (*models.User, error) {
	// Fetch the user by ID
	existingUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// Update fields with non-zero values from the updatedUser
	if updatedUser.Name != "" {
		existingUser.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		str, err := utils.HashPassword(updatedUser.Password)
		if err != nil {
			return nil, fmt.Errorf("error in updating password %v", err.Error())
		}
		existingUser.Password = str
	}

	// Call the repository to update the user
	if err := s.repo.UpdateUser(existingUser); err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return existingUser, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
