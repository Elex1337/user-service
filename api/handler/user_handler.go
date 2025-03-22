package handler

import (
	"github.com/Elex1337/user-service/internal/dto"
	"github.com/Elex1337/user-service/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler interface {
	CreateUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	GetUserByID(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

type UserHandlerImpl struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{userService}
}

// CreateUser godoc
// @Summary Создание юзера
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "userDto"
// @Success 201 {object} dto.UserResponseDTO
// @Router /users [post]
func (h *UserHandlerImpl) CreateUser(ctx echo.Context) error {
	var createDTO dto.CreateUserDTO

	if err := ctx.Bind(&createDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if createDTO.UserName == "" || createDTO.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Username and password are required",
		})
	}

	user, err := h.userService.CreateUser(createDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary Обновление юзера
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UpdateUserDTO true "userDto"
// @Success 200 {object} dto.UserResponseDTO
// @Router /users [put]
func (h *UserHandlerImpl) UpdateUser(ctx echo.Context) error {
	var updateDTO dto.UpdateUserDTO

	if err := ctx.Bind(&updateDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if updateDTO.UserName == "" || updateDTO.Password == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Username and password are required",
		})
	}

	user, err := h.userService.UpdateUser(updateDTO)
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, user)
}

// GetUserByID godoc
// @Summary получает юзера по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.UserResponseDTO
// @Router /users/{id} [get]
func (h *UserHandlerImpl) GetUserByID(ctx echo.Context) error {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Удаление пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandlerImpl) DeleteUser(ctx echo.Context) error {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User successfully deleted",
	})
}
