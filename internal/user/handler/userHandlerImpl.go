package handler

import (
	"net/http"

	"github.com/LuanTenorio/todo_go/internal/user/dto"
	usecase "github.com/LuanTenorio/todo_go/internal/user/useCase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type userHandlerImpl struct {
	UserUseCase usecase.UserUseCase
}

func NewUserUseCaseImpl(uc usecase.UserUseCase) UserHandler {
	return &userHandlerImpl{UserUseCase: uc}
}

func (h *userHandlerImpl) CreateUser(c echo.Context) error {
	userDto := new(dto.CreateUserDTO)

	if err := c.Bind(userDto); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return c.String(http.StatusBadRequest, "Bad request")
	}

	return c.JSON(http.StatusCreated, userDto)
}
