package handler

import (
	"log"
	"net/http"

	"github.com/LuanTenorio/todo_go/internal/requestError"
	"github.com/LuanTenorio/todo_go/internal/user/dto"
	usecase "github.com/LuanTenorio/todo_go/internal/user/useCase"
	"github.com/labstack/echo/v4"
)

const (
	errorOnCreateUser = "Error creating user"
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
		return c.JSON(http.StatusBadRequest, requestError.New(requestError.IncompatibleBody))
	} else if err := c.Validate(userDto); err != nil {
		return c.JSON(http.StatusBadRequest, requestError.New(err.Error()))
	}

	user, err := h.UserUseCase.CreateUser(c.Request().Context(), userDto)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusAccepted, requestError.New("internal server error"))
	}

	return c.JSON(http.StatusCreated, user)
}
