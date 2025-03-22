package service_test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/Elex1337/user-service/internal/dto"
	"github.com/Elex1337/user-service/internal/entity"
	"github.com/Elex1337/user-service/internal/repository/mocks"
	"github.com/Elex1337/user-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userService := service.NewUserService(mockRepo)

	userId := 1

	mockRepo.On("GetUserByID", userId).Return(entity.User{
		ID:        userId,
		UserName:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	mockRepo.On("DeleteUser", userId).Return(nil)

	err := userService.DeleteUser(userId)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userService := service.NewUserService(mockRepo)

	userId := 999

	mockRepo.On("GetUserByID", userId).Return(entity.User{}, errors.New("user not found"))

	err := userService.DeleteUser(userId)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	mockRepo.AssertExpectations(t)

	mockRepo.AssertNotCalled(t, "DeleteUser")
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userService := service.NewUserService(mockRepo)

	userId := 1
	createdAt := time.Now().Add(-24 * time.Hour)
	updatedAt := time.Now()

	updateDTO := dto.UpdateUserDTO{
		ID:       userId,
		UserName: "updateduser",
		Password: "newpassword123",
	}

	existingUser := entity.User{
		ID:        userId,
		UserName:  "testuser",
		Password:  "oldhashed",
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	updatedUser := entity.User{
		ID:        userId,
		UserName:  updateDTO.UserName,
		Password:  "newhashed",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	mockRepo.On("GetUserByID", userId).Return(existingUser, nil)
	mockRepo.On("UpdateUser", mock.Anything).Return(updatedUser, nil)

	result, err := userService.UpdateUser(updateDTO)

	assert.NoError(t, err)
	assert.Equal(t, updatedUser.ID, result.ID)
	assert.Equal(t, updatedUser.UserName, result.UserName)
	assert.Equal(t, updatedUser.CreatedAt, result.CreatedAt)
	assert.Equal(t, updatedUser.UpdatedAt, result.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userService := service.NewUserService(mockRepo)
	createDTO := dto.CreateUserDTO{
		UserName: "testuser",
		Password: "password123",
	}

	createdUser := entity.User{
		ID:        1,
		UserName:  createDTO.UserName,
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("CreateUser", mock.Anything).Return(createdUser, nil)

	result, err := userService.CreateUser(createDTO)

	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, result.ID)
	assert.Equal(t, createdUser.UserName, result.UserName)
	assert.Equal(t, createdUser.CreatedAt, result.CreatedAt)
	assert.Equal(t, createdUser.UpdatedAt, result.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	userService := service.NewUserService(mockRepo)

	userId := 1

	mockUser := entity.User{
		ID:        userId,
		UserName:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetUserByID", userId).Return(mockUser, nil)

	result, err := userService.GetUserByID(userId)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, result.ID)
	assert.Equal(t, mockUser.UserName, result.UserName)
	assert.Equal(t, mockUser.CreatedAt, result.CreatedAt)
	assert.Equal(t, mockUser.UpdatedAt, result.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
