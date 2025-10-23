package services

import (
	"errors"
	"fmt"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
	"restaurant-system/pkg/utils"
)

type UserServices interface {
	Registration(user *models.User) (*models.User, error)
	Login(email, password string) (*models.User, error)
	GetAll() ([]models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{userRepo: repo}
}

func (s *userService) Registration(user *models.User) (*models.User, error) {
	existing, _ := s.userRepo.GetByEmail(user.Email)
	if existing != nil {
		return nil, errors.New("user already exists with this email")
	}

	user.Password, _ = utils.HashPassword(user.Password)

	return user, s.userRepo.Create(user)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	userExit, err := s.userRepo.GetByEmail(email)

	if userExit == nil || err != nil {
		return nil, errors.New("internal server error or invalid email or password")
	}

	if !utils.CheckPassword(userExit.Password, password) {
		return nil, errors.New("wrong password")
	}

	return userExit, nil
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserById(id int) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)

	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *userService) UpdateUser(req_id int, updateUser *models.User) (*models.User, error) {
	updateUser.ID = req_id

	curUser, _ := s.userRepo.GetByID(updateUser.ID)

	if updateUser.Name != "" {
		curUser.Name = updateUser.Name
	}

	if updateUser.Email != "" {
		curUser.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		updateUser.Password, _ = utils.HashPassword(updateUser.Password)
		curUser.Password = updateUser.Password
	}

	if updateUser.PhoneNo != "" {
		curUser.PhoneNo = updateUser.PhoneNo
	}

	if updateUser.Role != "" {
		curUser.Role = updateUser.Role
	}

	if updateUser.ShiftTime != "" {
		curUser.ShiftTime = updateUser.ShiftTime
	}

	return curUser, s.userRepo.Update(curUser)
}

func (s *userService) DeleteUser(id int) error {
	user, _ := s.userRepo.GetByID(id)
	if user == nil {
		return errors.New("this user not Exit")
	}

	return s.userRepo.Delete(id)
}
