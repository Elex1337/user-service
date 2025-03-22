package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Elex1337/user-service/api/handler"
	"github.com/Elex1337/user-service/internal/dto"
	"github.com/Elex1337/user-service/internal/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupEcho() (*echo.Echo, *httptest.ResponseRecorder) {
	e := echo.New()
	rec := httptest.NewRecorder()
	return e, rec
}

func TestDeleteUser_Success(t *testing.T) {
	mockService := new(mocks.UserService)

	userHandler := handler.NewUserHandler(mockService)

	e, rec := setupEcho()

	userId := 1

	mockService.On("DeleteUser", userId).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := userHandler.DeleteUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Contains(t, response, "message")
	assert.Equal(t, "User successfully deleted", response["message"])

	mockService.AssertExpectations(t)
}

func TestDeleteUser_InvalidID(t *testing.T) {
	mockService := new(mocks.UserService)

	userHandler := handler.NewUserHandler(mockService)

	e, rec := setupEcho()

	req := httptest.NewRequest(http.MethodDelete, "/users/invalid", nil)

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := userHandler.DeleteUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Contains(t, response, "error")
	assert.Equal(t, "Invalid ID format", response["error"])

	mockService.AssertNotCalled(t, "DeleteUser")
}

func TestDeleteUser_UserNotFound(t *testing.T) {
	mockService := new(mocks.UserService)

	userHandler := handler.NewUserHandler(mockService)

	e, rec := setupEcho()

	userId := 999

	mockService.On("DeleteUser", userId).Return(errors.New("user not found"))

	req := httptest.NewRequest(http.MethodDelete, "/users/999", nil)

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := userHandler.DeleteUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Contains(t, response, "error")
	assert.Equal(t, "User not found", response["error"])

	mockService.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	mockService := new(mocks.UserService)

	userHandler := handler.NewUserHandler(mockService)

	e, rec := setupEcho()

	createDTO := dto.CreateUserDTO{
		UserName: "testuser",
		Password: "password123",
	}

	expectedResponse := dto.UserResponseDTO{
		ID:        1,
		UserName:  createDTO.UserName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockService.On("CreateUser", createDTO).Return(expectedResponse, nil)

	requestJSON, _ := json.Marshal(createDTO)
	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		strings.NewReader(string(requestJSON)),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	err := userHandler.CreateUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response dto.UserResponseDTO
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.UserName, response.UserName)

	mockService.AssertExpectations(t)
}
