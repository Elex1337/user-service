package service

import (
	"errors"
	"github.com/Elex1337/user-service/internal/dto"
	"github.com/Elex1337/user-service/internal/entity"
	"github.com/Elex1337/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(createDTO dto.CreateUserDTO) (dto.UserResponseDTO, error)
	GetUserByID(id int) (dto.UserResponseDTO, error)
	UpdateUser(updateDTO dto.UpdateUserDTO) (dto.UserResponseDTO, error)
	DeleteUser(id int) error
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo}
}

func (s *UserServiceImpl) CreateUser(createDTO dto.CreateUserDTO) (dto.UserResponseDTO, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponseDTO{}, errors.New("failed to hash password")
	}

	user := entity.User{
		UserName: createDTO.UserName,
		Password: string(hashedPassword),
	}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return dto.UserResponseDTO{
		ID:        createdUser.ID,
		UserName:  createdUser.UserName,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *UserServiceImpl) GetUserByID(id int) (dto.UserResponseDTO, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return dto.UserResponseDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserServiceImpl) UpdateUser(updateDTO dto.UpdateUserDTO) (dto.UserResponseDTO, error) {
	_, err := s.userRepo.GetUserByID(updateDTO.ID)
	if err != nil {
		return dto.UserResponseDTO{}, errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponseDTO{}, errors.New("failed to hash password")
	}

	user := entity.User{
		ID:       updateDTO.ID,
		UserName: updateDTO.UserName,
		Password: string(hashedPassword),
	}

	updatedUser, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return dto.UserResponseDTO{
		ID:        updatedUser.ID,
		UserName:  updatedUser.UserName,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}, nil
}
func (s *UserServiceImpl) DeleteUser(id int) error {
	_, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.userRepo.DeleteUser(id)
}
